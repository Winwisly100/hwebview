import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:hwebview/hwebview.dart';

void main() {
  const MethodChannel channel = MethodChannel('hwebview');

  TestWidgetsFlutterBinding.ensureInitialized();

  setUp(() {
    channel.setMockMethodCallHandler((MethodCall methodCall) async {
      return '42';
    });
  });

  tearDown(() {
    channel.setMockMethodCallHandler(null);
  });

  test('getPlatformVersion', () async {
    expect(await Hwebview.platformVersion, '42');
  });
}
