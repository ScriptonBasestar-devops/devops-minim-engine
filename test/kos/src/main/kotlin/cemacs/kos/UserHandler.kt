package cemacs.kos

import org.springframework.web.reactive.function.server.ServerRequest
import org.springframework.web.reactive.function.server.ServerResponse
import org.springframework.web.reactive.function.server.body
import org.springframework.web.reactive.function.server.bodyToServerSentEvents
import reactor.core.publisher.Flux
import java.time.Duration
import java.time.LocalDate
import java.time.format.DateTimeFormatter


@Suppress("UNUSED_PARAMETER")
class UserHandler {

	private val users = Flux.just(
			User("Foo", "Foo", LocalDate.now().minusDays(1)),
			User("Bar", "Bar", LocalDate.now().minusDays(10)),
			User("Baz", "Baz", LocalDate.now().minusDays(100)))

	private val userStream = Flux
			.zip(Flux.interval(Duration.ofMillis(100)), users.repeat())
			.map { it.t2 }

	fun findAll(req: ServerRequest) =
			ServerResponse.ok().body(users)

	fun findAllView(req: ServerRequest) =
			ServerResponse.ok().render("users", mapOf("users" to users.map { it.toDto() }))

	fun stream(req: ServerRequest) =
			ServerResponse.ok().bodyToServerSentEvents(userStream)

}

class UserDto(val firstName: String, val lastName: String, val birthDate: String)

fun User.toDto() = UserDto(firstName, lastName, birthDate.format(DateTimeFormatter.BASIC_ISO_DATE))
