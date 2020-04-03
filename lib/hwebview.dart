import 'dart:async';

import 'package:flutter/services.dart';

class Hwebview {
  static const MethodChannel _channel =
      const MethodChannel('hwebview');

  static Future<String> openWebview(String url) async {
    await _channel.invokeMethod('openWebview', url);
    return "";
  }
}

