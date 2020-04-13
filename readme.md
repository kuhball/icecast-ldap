# icecast-ldap
This container provides ldap authentication for icecast sources. 

Following env vars are used for the configuration:
- `ICECAST_AUTH_LDAP_SECURE` if true ldaps is used
- `ICECAST_AUTH_LDAP_SRV` hostname or IP from ldap server
- `ICECAST_AUTH_LDAP_DN` dn used for validation

Example icecast.xml mount part:
```xml
<mount>
    <mount-name>/*</mount-name>
    <authentication type="url">
        <option name="stream_auth" value="http://localhost:1337/"/>
    </authentication>
</mount>
```

Build container locally:
```shell script
podman build --rm -t scouball/icecast-ldap:latest .
```

Example kube file for podman pods:
[icecast.yml](icecast.yml)