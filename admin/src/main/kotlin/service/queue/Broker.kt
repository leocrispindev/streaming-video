package service.queue

import model.Document
import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.ProducerConfig
import org.apache.kafka.clients.producer.ProducerRecord
import java.util.Properties

object Broker {
    private var producer: KafkaProducer<String, Document>;

    init {
        val prop = Properties().also {
            it[ProducerConfig.BOOTSTRAP_SERVERS_CONFIG] = "http://localhost:9092"
            it[ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG] = DocumentSerializer::class
            it[ProducerConfig.RETRIES_CONFIG] = 2
        }

        producer = KafkaProducer(prop)
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
    fun sendSynchronous(topic: String, doc : Document) {
        val key = "video-id-" + doc.videoInfo.id

        producer.send(ProducerRecord( topic, key, doc)
        )
    }




}