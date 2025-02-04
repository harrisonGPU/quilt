//===-------- MergeCFuncGo.cpp - Transformations --------------------------===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//

#include "llvm/Transforms/Utils/MergeCFuncGo.h"
#include "llvm/ADT/SmallVector.h"
#include "llvm/IR/DerivedTypes.h"
#include "llvm/IR/Instructions.h"
#include "llvm/IR/Module.h"
#include "llvm/IR/PassManager.h"
#include "llvm/Support/raw_ostream.h"
#include <vector>

using namespace llvm;

static cl::opt<bool> RenameCallee_cs("rename-callee-cg", cl::init(false),
                                     cl::desc("rename the callee functions"));

#define DEBUG_TYPE "merge-c-func-go"

PreservedAnalyses MergeCFuncGoPass::run(Module &M, ModuleAnalysisManager &AM) {
  bool Changed = false;
  if (RenameCallee_cs) {
    RenameCallee(&M);
    Changed = true;
  }

  return Changed ? PreservedAnalyses::none() : PreservedAnalyses::all();
}

void MergeCFuncGoPass::RenameCallee(Module *M) {
  Function *mainFunc = M->getFunction("main.main");
  mainFunc->setName("main.callee");
}