# Disclaimer

This is not a serious project...

# Purpose
reads stdin
prints stdout

given a plist, prints json. given json prints a plist.

# Usage
json to plist
```
echo '{"foo":"bar"}' | ./jsonplistjson

<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>foo</key>
    <string>bar</string>
  </dict>
</plist>
```

plist to json

```
echo '<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd"\>
<plist version="1.0">
  <dict>
    <key>foo</key>
    <string>bar</string>
  </dict>
</plist>'|./jsonplistjson


{
  "foo": "bar"
}
```
