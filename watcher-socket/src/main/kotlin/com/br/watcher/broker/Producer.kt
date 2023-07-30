package com.br.watcher.broker

import com.br.watcher.model.VideoBrokerMessage
import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.ProducerConfig
import org.apache.kafka.clients.producer.ProducerRecord
import org.apache.kafka.common.serialization.StringSerializer
import java.util.*

object Producer {

    private var producer: KafkaProducer<String, VideoBrokerMessage>;

    init {
        val prop = Properties().also {
            it[ProducerConfig.BOOTSTRAP_SERVERS_CONFIG] = "http://localhost:9092"
            it[ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG] = "org.apache.kafka.common.serialization.StringSerializer"
            it[ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG] = "com.br.watcher.broker.DocumentSerializer"
            it[ProducerConfig.RETRIES_CONFIG] = 2
            it[ProducerConfig.TRANSACTIONAL_ID_CONFIG] = "producer-" + System.getProperty("admin.broker.producer.transaction.id")
        }

        producer = KafkaProducer(prop)
        producer.initTransactions()
        println("Producer OK")

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
    fun send(topic: String, doc : VideoBrokerMessage) {

        producer.send(ProducerRecord( topic, "", doc))
        println("Message send with success")
    }

}