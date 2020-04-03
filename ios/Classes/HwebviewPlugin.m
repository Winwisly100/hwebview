#import "HwebviewPlugin.h"
#if __has_include(<hwebview/hwebview-Swift.h>)
#import <hwebview/hwebview-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "hwebview-Swift.h"
#endif

@implementation HwebviewPlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftHwebviewPlugin registerWithRegistrar:registrar];
}
@end
