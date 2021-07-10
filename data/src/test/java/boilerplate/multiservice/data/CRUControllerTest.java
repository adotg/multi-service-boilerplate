package boilerplate.multiservice.data;

import lombok.SneakyThrows;
import org.junit.jupiter.api.*;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.HttpMethod;

import java.util.UUID;

@SpringBootTest
@TestInstance(TestInstance.Lifecycle.PER_CLASS)
public class CRUControllerTest extends AbstractTest {
    String appender = UUID.randomUUID().toString().split("-")[1];

    @Override
    @BeforeAll
    public void setUp() {
        super.setUp();
    }

    @SneakyThrows
    @Test
    @Order(1)
    public void insertingKeyShouldPass() {
        String value = sendRequest("/key/key_" + appender, HttpMethod.POST, "value_" + appender);
        Assertions.assertEquals("value_" + appender, value);
    }

    @SneakyThrows
    @Test
    @Order(2)
    public void gettingExistingKeyShouldPass() {
        String value = sendRequest("/key/key_" + appender, HttpMethod.GET, "");
        Assertions.assertEquals("value_" + appender, value);
    }
}
