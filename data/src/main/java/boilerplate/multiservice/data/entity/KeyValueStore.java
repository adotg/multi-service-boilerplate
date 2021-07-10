package boilerplate.multiservice.data.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor
@ToString
@Table(name="key_value_store")
public class KeyValueStore {
    @Id
    @GenericGenerator(name = "keygen", strategy = "boilerplate.multiservice.data.KeyIdentifierGenerator")
    @Column(name = "key_name")
    private String key;
    private String value;
}
