plugins {
    kotlin("jvm") version "2.1.21"
    id("org.jetbrains.intellij.platform") version "2.11.0"
}

group = "com.huajian"
version = "0.1.0"

repositories {
    mavenCentral()
    intellijPlatform {
        defaultRepositories()
    }
}

dependencies {
    intellijPlatform {
        val home = System.getProperty("user.home")
        val goideaHome = System.getenv("GOIDEA_HOME")?.trim()?.takeIf { it.isNotEmpty() }?.let(::File)
        val golandApp = listOf(
            goideaHome,
            File("$home/Applications/GoLand 2025.3.3.app"),
            File("/Applications/GoLand 2025.3.3.app"),
            File("$home/Applications/GoLand.app"),
            File("/Applications/GoLand.app")
        ).filterNotNull().firstOrNull { it.isDirectory }
            ?: error("GoLand.app not found. Set GOIDEA_HOME or install GoLand under ~/Applications or /Applications.")

        local(golandApp.absolutePath)
    }
}

intellijPlatform {
    pluginConfiguration {
        ideaVersion {
            sinceBuild = "253"
            untilBuild = "253.*"
        }
    }
}

java {
    toolchain {
        languageVersion.set(JavaLanguageVersion.of(21))
    }
}

kotlin {
    jvmToolchain(21)
}

tasks {
    named<org.gradle.api.tasks.bundling.Zip>("buildPlugin") {
        destinationDirectory.set(layout.projectDirectory)
    }

    withType<JavaCompile> {
        options.release.set(21)
    }

    withType<org.jetbrains.kotlin.gradle.tasks.KotlinCompile> {
        kotlinOptions.jvmTarget = "21"
    }
}
