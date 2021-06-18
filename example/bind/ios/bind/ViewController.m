// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import "ViewController.h"
@import Hello;  // Gomobile bind generated framework

@interface ViewController ()
@end

@implementation ViewController

@synthesize hjh;

- (void)loadView {
    [super loadView];
    hjh.text = HelloGreetings(@"emm iOS and Gopher");
    
}
@end
