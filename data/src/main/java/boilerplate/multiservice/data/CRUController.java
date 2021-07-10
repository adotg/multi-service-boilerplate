package boilerplate.multiservice.data;

import boilerplate.multiservice.data.entity.KeyValueStore;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cache.annotation.CachePut;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@Slf4j
public class CRUController {
    private KeyValueStoreRepo repo;

    @Autowired
    public void setRepo(KeyValueStoreRepo repo) {
        this.repo = repo;
    }

    @PostMapping("/key/{key_name}")
    @CachePut(value = "key_value_store", key="#key")
    public String upsertKeyValue(@PathVariable("key_name") String key, @RequestBody String value) {
        log.info("Insert/Update of key={} and value={} is requested", key, value);
        KeyValueStore keyValueStore = new KeyValueStore();
        keyValueStore.setKey(key);
        keyValueStore.setValue(value);
        repo.save(keyValueStore);
        return value;
    }

    @GetMapping("/key/{key_name}")
    @Cacheable(value = "key_value_store", key="#key")
    public String getKeyValue(@PathVariable("key_name") String key) {
        log.info("Retrieval of key={} is requested", key);
        KeyValueStore keyValueStoreByKey = repo.getKeyValueStoreByKey(key);
        if (keyValueStoreByKey == null) {
            return null;
        }
        return keyValueStoreByKey.getKey();
    }
}
