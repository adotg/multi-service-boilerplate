package boilerplate.multiservice.data;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.ApplicationRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cache.annotation.EnableCaching;
import org.springframework.context.annotation.Bean;
import org.springframework.core.env.Environment;

@SpringBootApplication
@Slf4j
@EnableCaching
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
            log.warn("Deploy Env = {}", System.getenv("DEPLOY_ENV"));
            log.warn("Service Name = {}", System.getenv("SERVICE_NAME"));

            String[] profiles = springEnv.getDefaultProfiles();
            for (String profile : profiles) log.warn("[Default Profiles] {}", profile);

            profiles = springEnv.getActiveProfiles();
            for (String profile : profiles) log.warn("[Active Profiles] {}", profile);
        };
    }
}
