package cemacs.kos

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class KosApplication

fun main(args: Array<String>) {
	runApplication<KosApplication>(*args)
}

