package com.br.watcher.broker

import org.apache.kafka.clients.consumer.ConsumerConfig
import org.apache.kafka.clients.consumer.ConsumerRecords
import org.apache.kafka.clients.consumer.KafkaConsumer
import java.util.Properties
import kotlin.time.Duration
import kotlin.time.Duration.Companion.seconds
import kotlin.time.toJavaDuration

class Consumer {

    private var consumer : KafkaConsumer<String, String>
    constructor(topic : String) {

        val prop = Properties().also {
            it[ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG] = ""
            it[ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG] = ""
            it[ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG] = ""
        }

        consumer = KafkaConsumer(prop)
        consumer.subscribe(listOf(topic))

    }

   fun StartPoll(time : Int) : ConsumerRecords<String, String> {
        return consumer.poll(time.seconds.toJavaDuration())
   }


}