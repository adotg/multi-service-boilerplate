package boilerplate.multiservice;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.ApplicationRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.core.env.Environment;

@SpringBootApplication
@Slf4j
public class App {
    private Environment springEnv;

    @Autowired
    public void setSpringEnv(Environment springEnv) {
        this.springEnv = springEnv;
    }

    public static void main( String[] args ) {
        SpringApplication.run(App.class);
    }

    @Bean
    ApplicationRunner applicationRunner()  {
        return args -> {
            String[] profiles = springEnv.getDefaultProfiles();
            for (String profile : profiles) log.warn("[Default Profiles] {}", profile);

            profiles = springEnv.getActiveProfiles();
            for (String profile : profiles) log.warn("[Active Profiles] {}", profile);
        };
    }
}
