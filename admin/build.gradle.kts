import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    kotlin("jvm") version "1.7.21"
    //id("com.jetbrains.exposed.gradle.plugin") version "0.2.1"
    application
}

group = "me.leonardo"
version = "1.0-SNAPSHOT"

repositories {
    //jcenter()
    mavenCentral()
    //maven("https://maven.pkg.jetbrains.space/public/p/kotlinx-html/maven")
}

dependencies {
    testImplementation(kotlin("test"))
    implementation("io.ktor:ktor-server-netty:1.6.0")
    implementation("io.ktor:ktor-html-builder:1.6.0")
    implementation("org.jetbrains.kotlinx:kotlinx-html-jvm:0.7.2")
    implementation("com.google.code.gson:gson:2.8.5")

    //Database
    implementation("org.jetbrains.exposed:exposed-core:0.40.1")
    implementation("org.jetbrains.exposed:exposed-dao:0.40.1")
    implementation("org.jetbrains.exposed:exposed-jdbc:0.40.1")
    implementation("mysql:mysql-connector-java:8.0.26")
    implementation("org.apache.kafka:kafka-clients:3.4.1")
    //implementation("gradle.plugin.com.jetbrains.exposed.gradle:plugin:0.2.1")
}

tasks.test {
    useJUnit()
}

tasks.withType<KotlinCompile>() {
    kotlinOptions.jvmTarget = "11"
}

//apply(plugin = "com.jetbrains.exposed.gradle.plugin")

application {
    mainClass.set("ServerKt")
}