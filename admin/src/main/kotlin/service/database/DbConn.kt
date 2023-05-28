package service.database

import org.jetbrains.exposed.sql.Database
import org.jetbrains.exposed.sql.transactions.transaction

object DbConn {
    private lateinit var conn : Database
    init {
      try {
         //conn = Database.connect("jdbc:mysql://localhost:3306/stream_content", driver = "org.h2.Driver",
              //user = "mysql", password = "mysql")

         conn = Database.connect("jdbc:mysql://localhost:3306/stream_content", user = "mysql", password = "mysql")
          transaction {
              //SchemaUtils.create(VideoInfoTb)
              println(conn.dialect.allTablesNames())
              //println(conn.config)
              println(message = "Connection ok")
          }


      }catch (e : Exception) {
        throw e;
      }
    }

     fun getDatabase(): Database {
        return conn;
    }
}