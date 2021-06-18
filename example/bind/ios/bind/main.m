// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import <UIKit/UIKit.h>
#import "AppDelegate.h"
@import Hello;  // Gomobile bind generated framework
int main(int argc, char * argv[]) {
    HelloRun();
    @autoreleasepool {
        return UIApplicationMain(argc, argv, nil, NSStringFromClass([AppDelegate class]));
    }
}
