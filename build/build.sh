#!/bin/zsh

VERSION="1.0.0" # 替换为你的版本号
ARTIFACT_ID="apiAuto" # 替换为你的应用名
GROUP_ID="com.hqy.auto" # 替换为你的组织ID

# 清理旧构建
rm -rf target
mkdir -p target/classes/{bin,meta}

# 构建多平台二进制
platforms=(
  "darwin amd64"
  "linux amd64"
  "windows amd64"
)

for platform in "${platforms[@]}"; do
  platform_split=(${=platform}) # zsh特有的数组分割语法
  GOOS=${platform_split[1]}
  GOARCH=${platform_split[2]}
  output_name=$ARTIFACT_ID

  if [ "$GOOS" = "windows" ]; then
    output_name+='.exe'
  fi

  echo "Building for $GOOS/$GOARCH..."
  env GOOS=$GOOS GOARCH=$GOARCH go build -o "target/classes/bin/${GOOS}-${GOARCH}/${output_name}" .
done

# 创建pom.xml
cat > target/pom.xml <<EOF
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <groupId>$GROUP_ID</groupId>
    <artifactId>$ARTIFACT_ID</artifactId>
    <version>$VERSION</version>
    <packaging>jar</packaging>
    <name>Go Application: $ARTIFACT_ID</name>
    <description>Go binary packaged for Maven</description>
</project>
EOF

# 打包JAR
echo "Creating JAR..."
jar cvf "target/${ARTIFACT_ID}-${VERSION}.jar" -C target/classes .
cp target/pom.xml "target/${ARTIFACT_ID}-${VERSION}.pom"

echo "Build complete. Files are in target/"
