# mvnrc
Maven remote repository version client

Basic authorization
```bash
mvnrc --username USER --password PASSWORD --url https://oss.sonatype.org/content/groups/public org.jboss.cdi.tck:cdi-tck-impl
```

Example release version
```bash
mvnrc --url https://oss.sonatype.org/content/groups/public org.jboss.cdi.tck:cdi-tck-impl

2.0.5.SP1
```

Example latest version
```bash
mvnrc --type latest --url https://oss.sonatype.org/content/groups/public org.jboss.cdi.tck:cdi-tck-impl

2.0.5.SP1
```

Example all versions in the remote repository:
```bash
mvnrc --type versions --url https://oss.sonatype.org/content/groups/public org.jboss.cdi.tck:cdi-tck-impl

2.0.1.Final
2.0.2-SNAPSHOT
2.0.2.Final
2.0.3-SNAPSHOT
2.0.3.Final
2.0.4-SNAPSHOT
2.0.4.Final
2.0.5-SNAPSHOT
2.0.5.Final
2.0.5.SP1

```