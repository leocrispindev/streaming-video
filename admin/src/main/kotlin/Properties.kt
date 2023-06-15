enum class Properties(private val key: String, private val value: String) {
    TOPIC_INDEX("admin.content.index.topic", "index-content");

    fun getValue(): String {
        return System.getProperty(key, value)
    }
}
