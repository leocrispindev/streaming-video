package com.br.watcher.broker

import org.apache.kafka.clients.admin.AdminClient
import org.apache.kafka.clients.admin.AdminClientConfig
import org.apache.kafka.clients.admin.NewTopic
import org.apache.kafka.common.KafkaException
import java.util.*
import kotlin.collections.ArrayList

object BrokerAdmin {

    // TODO define admin properties
    lateinit var admin : AdminClient

    init {
        val prop = Properties().also {
            it[AdminClientConfig.BOOTSTRAP_SERVERS_CONFIG] = "http://localhost:9092";
        }

        admin = AdminClient.create(prop)
        println("admin created");
    }

    // TODO create topic if not exists

     fun createTopic(name: String) : Boolean {
        try {
            val newTopic = NewTopic(name, 3, 0)

            val listTopics = mutableListOf(newTopic)
            val result = admin.createTopics(listTopics);

            println("Created topic: " + result.topicId(name));
            return true

        }catch (e : KafkaException) {
            throw e
        }
    }

     fun topicExist(topicName : String) : Boolean {
       return admin.listTopics().names().get().contains(topicName)
    }
}