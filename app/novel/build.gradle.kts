plugins {
    kotlin("jvm") version "2.3.20"
    id("org.jetbrains.intellij.platform") version "2.14.0"
}

group = "com.huajian"
version = "2026.1"

repositories {
    mavenCentral()
    intellijPlatform {
        defaultRepositories()
    }
}

dependencies {
    testImplementation(kotlin("test"))

    intellijPlatform {
        val home = System.getProperty("user.home")
        val goideaHome = System.getenv("GOIDEA_HOME")?.trim()?.takeIf { it.isNotEmpty() }?.let(::File)
        val golandApp = listOf(
            goideaHome,
            File("$home/Applications/GoLand 2026.1.app"),
            File("/Applications/GoLand 2026.1.app"),
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
            sinceBuild = "261"
            untilBuild = "261.*"
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

    withType<Test> {
        useJUnitPlatform()
    }

    withType<JavaCompile> {
        options.release.set(21)
    }

    withType<org.jetbrains.kotlin.gradle.tasks.KotlinJvmCompile> {
        compilerOptions {
            jvmTarget.set(org.jetbrains.kotlin.gradle.dsl.JvmTarget.JVM_21)
        }
    }
}
