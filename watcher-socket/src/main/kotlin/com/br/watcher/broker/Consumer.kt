package com.br.watcher.broker

import org.apache.kafka.clients.consumer.ConsumerConfig
import org.apache.kafka.clients.consumer.ConsumerRecords
import org.apache.kafka.clients.consumer.KafkaConsumer
import org.apache.kafka.common.serialization.StringDeserializer
import java.util.Properties
import kotlin.time.Duration
import kotlin.time.Duration.Companion.seconds
import kotlin.time.toJavaDuration

class Consumer {

    private var consumer : KafkaConsumer<String, String>
    constructor(topic : String, groupID: String) {

        val prop = Properties().also {
            it[ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG] = "http://localhost:9092"
            it[ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG] = "org.apache.kafka.common.serialization.StringDeserializer"
            it[ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG] = "org.apache.kafka.common.serialization.StringDeserializer"
            it[ConsumerConfig.GROUP_ID_CONFIG] = groupID
            it[ConsumerConfig.AUTO_OFFSET_RESET_CONFIG] = "earliest"
        }

        consumer = KafkaConsumer(prop)
        consumer.subscribe(listOf(topic))

    }

   fun StartPoll(time : Int) : ConsumerRecords<String, String> {

        return consumer.poll(time.seconds.toJavaDuration())
   }


}