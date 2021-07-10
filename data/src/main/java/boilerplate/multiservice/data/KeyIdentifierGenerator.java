package boilerplate.multiservice.data;

import org.hibernate.HibernateException;
import org.hibernate.engine.spi.SharedSessionContractImplementor;
import org.hibernate.id.IdentifierGenerator;
import org.hibernate.id.IdentityGenerator;

import java.io.Serializable;

public class KeyIdentifierGenerator extends IdentityGenerator implements IdentifierGenerator {
    @Override
    public Serializable generate(
        SharedSessionContractImplementor sess,
        Object o
    ) throws HibernateException {

        Serializable id = sess.getEntityPersister(null, o).getClassMetadata().getIdentifier(o, sess);
        return id != null ? id : super.generate(sess, o);
    }
}
