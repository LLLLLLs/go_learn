package stockcalculator

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"fyne.io/fyne"
)

const (
	dataDirName          = "data"
	usersDirName         = "users"
	stateFileName        = "portfolio_state.json"
	appStateFileName     = "app_state.json"
	stateFileMode        = 0o644
	dataDirectoryMode    = 0o755
	defaultUserID        = "default"
	defaultUserName      = "默认用户"
	legacyStateFileName  = "portfolio_state.json"
	createUserOptionName = "创建新用户"
)

type UserProfile struct {
	ID           string `json:"id"`
	Account      string `json:"account,omitempty"`
	Name         string `json:"name"`
	AvatarData   string `json:"avatar_data,omitempty"`
	PasswordSalt string `json:"password_salt,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
}

type appState struct {
	CurrentUserID string        `json:"current_user_id"`
	Users         []UserProfile `json:"users"`
}

type UserStore struct {
	State appState
}

func DataRootPath() string {
	if dataDir := strings.TrimSpace(os.Getenv("STOCKCALC_DATA_DIR")); dataDir != "" {
		return dataDir
	}
	return filepath.Join(".", dataDirName)
}

func AppStatePath() string {
	return filepath.Join(DataRootPath(), appStateFileName)
}

func LegacyPortfolioStatePath() string {
	return filepath.Join(DataRootPath(), legacyStateFileName)
}

func UserPortfolioStatePath(userID string) string {
	return filepath.Join(DataRootPath(), usersDirName, userID, stateFileName)
}

func LoadUserStore() (*UserStore, error) {
	store := &UserStore{
		State: appState{
			CurrentUserID: defaultUserID,
			Users: []UserProfile{
				{ID: defaultUserID, Name: defaultUserName},
			},
		},
	}

	raw, err := os.ReadFile(AppStatePath())
	if err == nil {
		if err := json.Unmarshal(raw, &store.State); err != nil {
			return nil, fmt.Errorf("读取用户配置失败: %w", err)
		}
	}
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("读取用户配置失败: %w", err)
	}

	store.normalize()
	if err := os.MkdirAll(filepath.Join(DataRootPath(), usersDirName), dataDirectoryMode); err != nil {
		return nil, fmt.Errorf("创建用户数据目录失败: %w", err)
	}
	if err := store.migrateLegacyState(); err != nil {
		return nil, err
	}
	if err := store.Save(); err != nil {
		return nil, err
	}
	return store, nil
}

func LoadPortfolioForUser(store *UserStore, prefs fyne.Preferences) (*Portfolio, string, error) {
	profile := store.CurrentUser()
	statePath := UserPortfolioStatePath(profile.ID)
	portfolio, err := LoadPortfolioFromStorage(statePath, prefs)
	if err != nil {
		return nil, "", err
	}
	return portfolio, statePath, nil
}

func LoadPortfolioFileForUser(store *UserStore) (*Portfolio, string, error) {
	profile := store.CurrentUser()
	statePath := UserPortfolioStatePath(profile.ID)
	portfolio, err := LoadPortfolioFromStorage(statePath, nil)
	if err != nil {
		return nil, "", err
	}
	return portfolio, statePath, nil
}

func LoadPortfolioFromStorage(statePath string, prefs fyne.Preferences) (*Portfolio, error) {
	raw, err := os.ReadFile(statePath)
	if err == nil {
		return LoadPortfolio(string(raw))
	}
	if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("读取本地数据文件失败: %w", err)
	}

	if prefs == nil {
		return LoadPortfolio("")
	}

	legacyRaw := strings.TrimSpace(prefs.String(statePreferenceKey))
	portfolio, err := LoadPortfolio(legacyRaw)
	if err != nil {
		return nil, err
	}
	if legacyRaw == "" {
		return portfolio, nil
	}
	if err := portfolio.SaveToFile(statePath); err != nil {
		return nil, err
	}
	prefs.RemoveValue(statePreferenceKey)
	return portfolio, nil
}

func (p *Portfolio) SaveToFile(statePath string) error {
	payload, err := p.MarshalState()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(statePath), dataDirectoryMode); err != nil {
		return fmt.Errorf("创建数据目录失败: %w", err)
	}
	if err := os.WriteFile(statePath, payload, stateFileMode); err != nil {
		return fmt.Errorf("写入本地数据文件失败: %w", err)
	}
	return nil
}

func (s *UserStore) Save() error {
	s.normalize()
	payload, err := json.MarshalIndent(s.State, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化用户配置失败: %w", err)
	}
	if err := os.MkdirAll(DataRootPath(), dataDirectoryMode); err != nil {
		return fmt.Errorf("创建数据目录失败: %w", err)
	}
	if err := os.WriteFile(AppStatePath(), payload, stateFileMode); err != nil {
		return fmt.Errorf("写入用户配置失败: %w", err)
	}
	return nil
}

func (s *UserStore) normalize() {
	if len(s.State.Users) == 0 {
		s.State.Users = []UserProfile{{ID: defaultUserID, Name: defaultUserName}}
	}

	seen := make(map[string]bool, len(s.State.Users))
	users := make([]UserProfile, 0, len(s.State.Users))
	for _, user := range s.State.Users {
		user.ID = strings.TrimSpace(user.ID)
		user.Name = strings.TrimSpace(user.Name)
		if user.ID == "" || user.Name == "" || seen[user.ID] {
			continue
		}
		seen[user.ID] = true
		users = append(users, user)
	}
	if len(users) == 0 {
		users = []UserProfile{{ID: defaultUserID, Name: defaultUserName}}
	}
	sort.Slice(users, func(i, j int) bool {
		if users[i].ID == defaultUserID {
			return true
		}
		if users[j].ID == defaultUserID {
			return false
		}
		return users[i].Name < users[j].Name
	})
	s.State.Users = users

	if s.findUserByID(s.State.CurrentUserID) == nil {
		s.State.CurrentUserID = s.State.Users[0].ID
	}
}

func (s *UserStore) migrateLegacyState() error {
	legacyPath := LegacyPortfolioStatePath()
	if _, err := os.Stat(legacyPath); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	defaultPath := UserPortfolioStatePath(defaultUserID)
	if _, err := os.Stat(defaultPath); err == nil {
		return os.Remove(legacyPath)
	}

	raw, err := os.ReadFile(legacyPath)
	if err != nil {
		return fmt.Errorf("读取旧版数据文件失败: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(defaultPath), dataDirectoryMode); err != nil {
		return fmt.Errorf("创建默认用户目录失败: %w", err)
	}
	if err := os.WriteFile(defaultPath, raw, stateFileMode); err != nil {
		return fmt.Errorf("迁移旧版数据文件失败: %w", err)
	}
	if err := os.Remove(legacyPath); err != nil {
		return fmt.Errorf("删除旧版数据文件失败: %w", err)
	}
	return nil
}

func (s *UserStore) CurrentUser() UserProfile {
	s.normalize()
	user := s.findUserByID(s.State.CurrentUserID)
	if user == nil {
		return s.State.Users[0]
	}
	return *user
}

func (s *UserStore) Users() []UserProfile {
	s.normalize()
	users := make([]UserProfile, len(s.State.Users))
	copy(users, s.State.Users)
	return users
}

func (s *UserStore) UserNames() []string {
	users := s.Users()
	names := make([]string, 0, len(users))
	for _, user := range users {
		names = append(names, user.Name)
	}
	return names
}

func (s *UserStore) UserOptions() []string {
	options := append([]string(nil), s.UserNames()...)
	return append(options, createUserOptionName)
}

func (s *UserStore) SetCurrentUserByName(name string) (UserProfile, error) {
	name = strings.TrimSpace(name)
	for _, user := range s.State.Users {
		if user.Name == name {
			s.State.CurrentUserID = user.ID
			return user, s.Save()
		}
	}
	return UserProfile{}, errors.New("用户不存在")
}

func (s *UserStore) SetCurrentUserByID(id string) (UserProfile, error) {
	id = strings.TrimSpace(id)
	user := s.findUserByID(id)
	if user == nil {
		return UserProfile{}, errors.New("用户不存在")
	}
	s.State.CurrentUserID = user.ID
	return *user, s.Save()
}

func (s *UserStore) CreateUser(name string) (UserProfile, error) {
	var err error
	name, err = s.validateUniqueUserName("", name)
	if err != nil {
		return UserProfile{}, err
	}

	id := slugifyUserID(name)
	if id == "" {
		id = "user-" + time.Now().Format("20060102150405")
	}

	originalID := id
	index := 2
	for s.findUserByID(id) != nil {
		id = fmt.Sprintf("%s-%d", originalID, index)
		index++
	}

	user := UserProfile{ID: id, Name: name}
	s.State.Users = append(s.State.Users, user)
	s.State.CurrentUserID = user.ID
	if err := s.Save(); err != nil {
		return UserProfile{}, err
	}
	return user, nil
}

func (s *UserStore) RegisterUser(account, password, passwordConfirm, name string) (UserProfile, error) {
	account, err := validateAccount(account)
	if err != nil {
		return UserProfile{}, err
	}
	if err := validatePassword(password); err != nil {
		return UserProfile{}, err
	}
	if password != passwordConfirm {
		return UserProfile{}, errors.New("两次输入的密码不一致")
	}
	name, err = s.validateUniqueUserName("", name)
	if err != nil {
		return UserProfile{}, err
	}
	if s.findUserByAccount(account) != nil {
		return UserProfile{}, errors.New("账号已存在")
	}

	id := slugifyUserID(account)
	if id == "" {
		id = "user-" + time.Now().Format("20060102150405")
	}
	originalID := id
	index := 2
	for s.findUserByID(id) != nil {
		id = fmt.Sprintf("%s-%d", originalID, index)
		index++
	}

	salt, hash, err := hashPassword(password)
	if err != nil {
		return UserProfile{}, err
	}
	user := UserProfile{ID: id, Account: account, Name: name, PasswordSalt: salt, PasswordHash: hash}
	s.State.Users = append(s.State.Users, user)
	s.State.CurrentUserID = user.ID
	if err := s.Save(); err != nil {
		return UserProfile{}, err
	}
	return user, nil
}

func (s *UserStore) Authenticate(account, password string) (UserProfile, error) {
	account, err := validateAccount(account)
	if err != nil {
		return UserProfile{}, errors.New("账号或密码不正确")
	}
	user := s.findUserByAccount(account)
	if user == nil || user.PasswordHash == "" || user.PasswordSalt == "" || !verifyPassword(password, user.PasswordSalt, user.PasswordHash) {
		return UserProfile{}, errors.New("账号或密码不正确")
	}
	return *user, nil
}

func (s *UserStore) UpdateUserProfile(id, name, avatarData string, resetAvatar bool, currentPassword, newPassword, newPasswordConfirm string) (UserProfile, error) {
	name, err := s.validateUniqueUserName(id, name)
	if err != nil {
		return UserProfile{}, err
	}

	user := s.findUserByID(id)
	if user == nil {
		return UserProfile{}, errors.New("用户不存在")
	}
	user.Name = name
	if resetAvatar {
		user.AvatarData = ""
	} else if avatarData != "" {
		user.AvatarData = avatarData
	}

	changingPassword := strings.TrimSpace(currentPassword) != "" || strings.TrimSpace(newPassword) != "" || strings.TrimSpace(newPasswordConfirm) != ""
	if changingPassword {
		if strings.TrimSpace(currentPassword) == "" {
			return UserProfile{}, errors.New("请输入当前密码")
		}
		if user.PasswordHash == "" || user.PasswordSalt == "" || !verifyPassword(currentPassword, user.PasswordSalt, user.PasswordHash) {
			return UserProfile{}, errors.New("当前密码不正确")
		}
		if err := validatePassword(newPassword); err != nil {
			return UserProfile{}, err
		}
		if newPassword != newPasswordConfirm {
			return UserProfile{}, errors.New("两次输入的新密码不一致")
		}
		salt, hash, err := hashPassword(newPassword)
		if err != nil {
			return UserProfile{}, err
		}
		user.PasswordSalt = salt
		user.PasswordHash = hash
	}

	if err := s.Save(); err != nil {
		return UserProfile{}, err
	}
	return *user, nil
}

func (s *UserStore) RenameUser(id, name string) (UserProfile, error) {
	name, err := s.validateUniqueUserName(id, name)
	if err != nil {
		return UserProfile{}, err
	}

	user := s.findUserByID(id)
	if user == nil {
		return UserProfile{}, errors.New("用户不存在")
	}
	user.Name = name
	if err := s.Save(); err != nil {
		return UserProfile{}, err
	}
	return *user, nil
}

func (s *UserStore) findUserByID(id string) *UserProfile {
	for i := range s.State.Users {
		if s.State.Users[i].ID == id {
			return &s.State.Users[i]
		}
	}
	return nil
}

func (s *UserStore) findUserByAccount(account string) *UserProfile {
	account = strings.ToLower(strings.TrimSpace(account))
	for i := range s.State.Users {
		if strings.EqualFold(s.State.Users[i].Account, account) {
			return &s.State.Users[i]
		}
	}
	return nil
}

func (s *UserStore) validateUniqueUserName(exceptID, name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", errors.New("用户名不能为空")
	}
	for _, user := range s.State.Users {
		if user.ID != exceptID && strings.EqualFold(user.Name, name) {
			return "", errors.New("用户名不能重复")
		}
	}
	return name, nil
}

func validateAccount(account string) (string, error) {
	account = strings.ToLower(strings.TrimSpace(account))
	if account == "" {
		return "", errors.New("账号不能为空")
	}
	for _, r := range account {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' || r == '.' {
			continue
		}
		return "", errors.New("账号只能包含字母、数字、横线、下划线和点")
	}
	return account, nil
}

func validatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("密码不能为空")
	}
	return nil
}

func hashPassword(password string) (string, string, error) {
	saltBytes := make([]byte, 16)
	if _, err := rand.Read(saltBytes); err != nil {
		return "", "", fmt.Errorf("生成密码盐失败: %w", err)
	}
	salt := hex.EncodeToString(saltBytes)
	sum := sha256.Sum256([]byte(salt + ":" + password))
	return salt, hex.EncodeToString(sum[:]), nil
}

func verifyPassword(password, salt, expectedHash string) bool {
	sum := sha256.Sum256([]byte(salt + ":" + password))
	actualHash := hex.EncodeToString(sum[:])
	return subtle.ConstantTimeCompare([]byte(actualHash), []byte(expectedHash)) == 1
}

func slugifyUserID(name string) string {
	var builder strings.Builder
	lastDash := false

	for _, r := range strings.ToLower(strings.TrimSpace(name)) {
		switch {
		case r >= 'a' && r <= 'z':
			builder.WriteRune(r)
			lastDash = false
		case r >= '0' && r <= '9':
			builder.WriteRune(r)
			lastDash = false
		case r == '-' || r == '_' || r == ' ':
			if builder.Len() > 0 && !lastDash {
				builder.WriteByte('-')
				lastDash = true
			}
		}
	}

	return strings.Trim(builder.String(), "-")
}
