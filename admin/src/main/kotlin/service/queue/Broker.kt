package service.queue

import model.Document
import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.ProducerConfig
import org.apache.kafka.clients.producer.ProducerRecord
import org.apache.kafka.common.serialization.StringSerializer
import java.util.Properties
import java.util.UUID

object Broker {
    private var producer: KafkaProducer<String, Document>;

    init {
        val prop = Properties().also {
            it[ProducerConfig.BOOTSTRAP_SERVERS_CONFIG] = "http://localhost:9092"
            it[ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG] = "org.apache.kafka.common.serialization.StringSerializer"
            it[ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG] = "service.queue.DocumentSerializer"
            it[ProducerConfig.RETRIES_CONFIG] = 2
            it[ProducerConfig.TRANSACTIONAL_ID_CONFIG] = "producer-" + System.getProperty("admin.broker.producer.transaction.id")
        }

        producer = KafkaProducer(prop)
        producer.initTransactions()
        println("Broker OK")

    }

    fun initTransaction() {
        producer.initTransactions()
    }
    fun beginTransaction() {
        producer.beginTransaction()
    }

    fun abortTransaction() {
        producer.abortTransaction()
    }

    fun commit() {
        producer.commitTransaction()
    }
    fun send(topic: String, doc : Document) {

        producer.send(ProducerRecord( topic, doc.key, doc))
    }




}