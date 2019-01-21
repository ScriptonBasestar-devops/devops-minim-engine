package cemacs.kos

import com.samskivert.mustache.Mustache
import org.springframework.context.MessageSource
import org.springframework.core.io.ClassPathResource
import org.springframework.http.MediaType.TEXT_HTML
import org.springframework.web.reactive.function.server.RenderingResponse
import org.springframework.web.reactive.function.server.router
import reactor.core.publisher.toMono
import java.util.*

class Routes(private val userHandler: UserHandler,
			 private val messageSource: MessageSource) {

	fun router() = router {
		accept(TEXT_HTML).nest {
			GET("/") { ok().render("home") }
		}
//		"/api".nest {
//			accept(APPLICATION_JSON).nest {
//				GET("/users", userHandler::findAll)
//			}
//
//		}
		resources("/**", ClassPathResource("static/"))
	}.filter { request, next ->
		next.handle(request).flatMap {
			if (it is RenderingResponse) RenderingResponse.from(it).modelAttributes(attributes(Locale.KOREA, messageSource)).build() else it.toMono()
		}
	}

	private fun attributes(locale: Locale, messageSource: MessageSource) = mutableMapOf<String, Any>(
			"i18n" to Mustache.Lambda { frag, out ->
				val tokens = frag.execute().split("|")
				out.write(messageSource.getMessage(tokens[0], tokens.slice(IntRange(1, tokens.size - 1)).toTypedArray(), locale)) })
}
