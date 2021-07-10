package boilerplate.multiservice.data;

import boilerplate.multiservice.data.entity.KeyValueStore;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface KeyValueStoreRepo extends CrudRepository<KeyValueStore, String> {
    KeyValueStore getKeyValueStoreByKey(String key);
}
