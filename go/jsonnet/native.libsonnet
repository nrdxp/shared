{
  getConfig(): std.native('getConfig')(),
  getEnv(key): std.native('getEnv')(key),
  getPath(path, fallback=null): std.native('getPath')(path, fallback),
  getRecord(type, name, fallback=null): std.native('getRecord')(type, name, fallback),
  regexMatch(regex, string): std.native('regexMatch')(regex, string),
}
