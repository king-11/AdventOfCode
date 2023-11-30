plugins {
    kotlin("jvm") version "1.9.21"
}

sourceSets {
    main {
        kotlin.srcDir("src/")
        java.srcDir("src/")
    }
}

tasks {
    wrapper {
        gradleVersion = "8.5"
    }
}