/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/osystem.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0035() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| osystem.cl                                                  |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
//-------------------------------------------------------------------
// This file contains the gloabal parameter objects and the key methods
// This are the key methods
//
//  c_type(x)  is the CLAIRE type of x
//  c_code(x)  is an optimized instruction
//  & c_code(x,s) is an optimized expression of sort s
//  c_sort(x)  the sort of the expression x (precise sort)
//  g_throw(x) boolean that says if x may throw an exception
//
// the sorts are integer, float, char, object, EID
//-----------------------------------------------------------------
// ******************************************************************
// *   Table of contents                                            *
// *    Part 1: General Global Variables and Properties             *
// *    Part 2: The defaults for c_type, c_code, c_gc and c_sort    *
// *    Part 3: g_throw and status(m:method)                        *
// *    Part 4: Names & identifiers management                      *
// ******************************************************************
//
// import
// Compile/index :: Kernel/index   (1) should not be needed (inherited by iClaire)
// Compile/typing :: Kernel/typing
// where to find the CLAIRE libraries
/* {1} OPT.The go function for: home(_CL_obj:void) [] */
func F_home_void () *ClaireString  { 
    // use function body compiling 
return  F_getenv_string(MakeString("CLAIRE_HOME"))
    } 
  
// The EID go function for: home @ void (throw: false) 
func E_home_void (_CL_obj EID) EID { 
    return EID{/*(sm for home @ void= string)*/ F_home_void( ).Id(),0}} 
  
// TO CHANGE -> simply read claire_home
// CLAIRE4 uses EID where CLAIRE used C++ OID (integer representation)
// used as a marker for form EID
// ******************************************************************
// *    Part 1: General Global Variables and Properties             *
// ******************************************************************
// we use an optimizer object with all the necessary resources
// they are all private.
// update on strings   v3.3.46
// The meta_compiler contains the definition of the compiler flags and slots
// that are important for the user. Other stuff is hidden in OPT
// v3.2.56: record -O option
// code producer are defined in Generate
// but the stub is define in Optimize to have access to current_file
// name of the file being compiled
// we use a global variable to hide the indirection through the producer
// this is kept in CLAIRE 4.0 so that the C++ compiler could be re-introduced
// new in CLAIRE4: create an automated comment
// the three variables that are used in the main files
// safety:
//       0  -> super-safe (check returns & gc safe)
//       1  -> safe
//       2  -> we trust explicit types & super
//       3  -> no overflow checking (integer & arrays)
//       4  -> no selector errors, no range error
//       5  -> cross-compiling (i.e. no errors)
//       6  -> unsafe (no GC)
// Id(compiler.options))
// re-definable items for bootstrap modifications
// Compile/make_float_function :: property(Core/open = 3)
// Compile/c_expression :: property(Core/open = 3)
// other useful properties shared between Optimize & Generate
// Optimizer version of sorts
// code with strict (stupid) type
// new: allow future overload !!
// compiler instantiation
// fast instantiation if all any slots are known
// how to compile a type expression
// these are the classes defined especially for this module
// Compile/to_CL <: Optimized_instruction(arg:any,set_arg:class)
// Compile/to_C <: Optimized_instruction(arg:any,set_arg:class)
// was to_C()
// Patterns are calls p(X) that are seen as a type expression
// the tuple is made into a list
// OPT contains all the parameters for the optimizer
// pragma for the compiler  => MOVED TO LANGUAGE in CLAIRE 4
// this pragma tells to compile with full safety (include arithmetic checks)
/* {1} OPT.The go function for: safe(x:any) [] */
func F_safe_any (x *ClaireAny ) *ClaireAny  { 
    // use function body compiling 
return  x
    } 
  
// The EID go function for: safe @ any (throw: false) 
func E_safe_any (x EID) EID { 
    return /*(sm for safe @ any= any)*/ F_safe_any(ANY(x) ).ToEID()} 
  
/* {1} OPT.The go function for: claire_safe_any_type */
func F_claire_safe_any_type (x *ClaireType ) EID { 
    /* eid body: x */
    var Result EID 
    Result = EID{x.Id(),0}
    return Result} 
  
  
// The dual EID go function for: "claire_safe_any_type" 
func E_claire_safe_any_type (x EID) EID { 
    return F_claire_safe_any_type(ToType(OBJ(x)))} 
  
// ******************************************************************
// *    Part 2: The defaults for c_type, c_code and c_sort          *
// ******************************************************************
// basic type inference
/* {1} OPT.The go function for: c_type(self:any) [] */
func F_c_type_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0036 *ClaireVariable   = To_Variable(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var r *ClaireAny   = Core.F_get_property(C_range,ToObject(g0036.Id()))
          /* noccur = 6 */
          if ((r == CNULL) || 
              (r == C_EID.Id())) /* If:5 */{ 
            Result = EID{C_any.Id(),0}
            } else {
            var g0049I *ClaireBoolean  
            if (r.Isa.IsIn(C_Union) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0037 *ClaireUnion   = To_Union(r)
                /* noccur = 1 */
                g0049I = Equal(g0037.T1.Id(),CEMPTY.Id())
                /* Let-7 */} 
              } else {
              g0049I = CFALSE
              /* If-6 */} 
            if (g0049I == CTRUE) /* If:6 */{ 
              Result = EID{To_Union(To_Union(r).T2.Id()).T2.Id(),0}
              } else {
              Result = EID{F_Optimize_ptype_type(ToType(r)).Id(),0}
              /* If-6 */} 
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0038 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var r *ClaireType   = g0038.Range
          /* noccur = 2 */
          if (F_boolean_I_any(r.Id()) == CTRUE) /* If:5 */{ 
            Result = EID{r.Id(),0}
            } else {
            Result = EID{MakeConstantSet(g0038.Value).Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0039 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        /* noccur = 1 */
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[215] the symbol ~A is unbound_symbol").Id(),0},EID{g0039.Name.Id(),0}))
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_error) == CTRUE) /* If:2 */{ 
      Result = EID{CEMPTY.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Update) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0041 *Language.Update   = Language.To_Update(self)
        /* noccur = 1 */
        Result = Core.F_CALL(C_c_type,ARGS(g0041.Value.ToEID()))
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Construct) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0042 *Language.Construct   = Language.To_Construct(self)
        /* noccur = 4 */
        if ((g0042.Isa.IsIn(Language.C_List) != CTRUE) && 
            (g0042.Isa.IsIn(Language.C_Set) != CTRUE)) /* If:4 */{ 
          Result = EID{C_any.Id(),0}
          } else {
          /* Let:5 */{ 
            var _Zres *ClaireType   = ToType(CEMPTY.Id())
            /* noccur = 5 */
            /* For:6 */{ 
              var _Zx *ClaireAny  
              _ = _Zx
              Result= EID{CFALSE.Id(),0}
              var _Zx_support *ClaireList  
              _Zx_support = g0042.Args
              _Zx_len := _Zx_support.Length()
              for i_it := 0; i_it < _Zx_len; i_it++ { 
                _Zx = _Zx_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                if (F_boolean_I_any(_Zres.Id()) == CTRUE) /* If:8 */{ 
                  var _Zres_try00509 EID 
                  /* Let:9 */{ 
                    var g0051UU *ClaireClass  
                    /* noccur = 1 */
                    var g0051UU_try005210 EID 
                    /* Let:10 */{ 
                      var g0053UU *ClaireType  
                      /* noccur = 1 */
                      var g0053UU_try005411 EID 
                      g0053UU_try005411 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                      /* ERROR PROTECTION INSERTED (g0053UU-g0051UU_try005210) */
                      if ErrorIn(g0053UU_try005411) {g0051UU_try005210 = g0053UU_try005411
                      } else {
                      g0053UU = ToType(OBJ(g0053UU_try005411))
                      g0051UU_try005210 = EID{g0053UU.Class_I().Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0051UU-_Zres_try00509) */
                    if ErrorIn(g0051UU_try005210) {_Zres_try00509 = g0051UU_try005210
                    } else {
                    g0051UU = ToClass(OBJ(g0051UU_try005210))
                    _Zres_try00509 = EID{Core.F_meet_class(ToClass(_Zres.Id()),g0051UU).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (_Zres-void_try8) */
                  if ErrorIn(_Zres_try00509) {void_try8 = _Zres_try00509
                  } else {
                  _Zres = ToType(OBJ(_Zres_try00509))
                  void_try8 = EID{_Zres.Id(),0}
                  }
                  } else {
                  var _Zres_try00559 EID 
                  /* Let:9 */{ 
                    var g0056UU *ClaireType  
                    /* noccur = 1 */
                    var g0056UU_try005710 EID 
                    g0056UU_try005710 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                    /* ERROR PROTECTION INSERTED (g0056UU-_Zres_try00559) */
                    if ErrorIn(g0056UU_try005710) {_Zres_try00559 = g0056UU_try005710
                    } else {
                    g0056UU = ToType(OBJ(g0056UU_try005710))
                    _Zres_try00559 = EID{g0056UU.Class_I().Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (_Zres-void_try8) */
                  if ErrorIn(_Zres_try00559) {void_try8 = _Zres_try00559
                  } else {
                  _Zres = ToType(OBJ(_Zres_try00559))
                  void_try8 = EID{_Zres.Id(),0}
                  }
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-Result) */
                if ErrorIn(void_try8) {Result = void_try8
                Result = void_try8
                break
                } else {
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* Let:6 */{ 
              var g0058UU *ClaireClass  
              /* noccur = 1 */
              if (g0042.Isa.IsIn(Language.C_Set) == CTRUE) /* If:7 */{ 
                g0058UU = C_set
                } else {
                g0058UU = C_list
                /* If-7 */} 
              Result = EID{Core.F_nth_class1(g0058UU,_Zres).Id(),0}
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0045 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 1 */
        Result = ToException(Core.C_general_error.Make(MakeString("c_type of ~S is not defined").Id(),MakeConstantList(g0045.Id().Isa.Id()).Id())).Close()
        /* Let-3 */} 
      } else {
      Result = EID{MakeConstantSet(self).Id(),0}
      /* If-2 */} 
    return RangeCheck(ToType(C_type.Id()),Result)} 
  
// The EID go function for: c_type @ any (throw: true) 
func E_c_type_any (self EID) EID { 
    return /*(sm for c_type @ any= EID)*/ F_c_type_any(ANY(self) )} 
  
// compile into a sort and checks strict type matching (naive/stupid)
/* {1} OPT.The go function for: Compile/c_strict_code(x:any,s:class) [] */
func F_Compile_c_strict_code_any (x *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0059UU *ClaireAny  
      /* noccur = 1 */
      var g0059UU_try00603 EID 
      g0059UU_try00603 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{s.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0059UU-Result) */
      if ErrorIn(g0059UU_try00603) {Result = g0059UU_try00603
      } else {
      g0059UU = ANY(g0059UU_try00603)
      Result = Core.F_CALL(C_Compile_c_strict_check,ARGS(g0059UU.ToEID(),EID{s.Id(),0}))
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/c_strict_code @ any (throw: true) 
func E_Compile_c_strict_code_any (x EID,s EID) EID { 
    return /*(sm for Compile/c_strict_code @ any= EID)*/ F_Compile_c_strict_code_any(ANY(x),ToClass(OBJ(s)) )} 
  
// CLAIRE 4: introduce C_cast so that psort(x) is what is expected (s)
/* {1} OPT.The go function for: Compile/c_strict_check(x:any,s:class) [] */
func F_Compile_c_strict_check_any (x *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    var g0061I *ClaireBoolean  
    var g0061I_try00622 EID 
    /* and:2 */{ 
      var v_and2 *ClaireBoolean  
      
      v_and2 = s.IsIn(C_object)
      if (v_and2 == CFALSE) {g0061I_try00622 = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and2_try00634 EID 
        /* Let:4 */{ 
          var g0064UU *ClaireBoolean  
          /* noccur = 1 */
          var g0064UU_try00655 EID 
          /* Let:5 */{ 
            var g0066UU *ClaireClass  
            /* noccur = 1 */
            var g0066UU_try00676 EID 
            g0066UU_try00676 = Language.F_static_type_any(x)
            /* ERROR PROTECTION INSERTED (g0066UU-g0064UU_try00655) */
            if ErrorIn(g0066UU_try00676) {g0064UU_try00655 = g0066UU_try00676
            } else {
            g0066UU = ToClass(OBJ(g0066UU_try00676))
            g0064UU_try00655 = EID{g0066UU.IsIn(s).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0064UU-v_and2_try00634) */
          if ErrorIn(g0064UU_try00655) {v_and2_try00634 = g0064UU_try00655
          } else {
          g0064UU = ToBoolean(OBJ(g0064UU_try00655))
          v_and2_try00634 = EID{g0064UU.Not.Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and2-g0061I_try00622) */
        if ErrorIn(v_and2_try00634) {g0061I_try00622 = v_and2_try00634
        } else {
        v_and2 = ToBoolean(OBJ(v_and2_try00634))
        if (v_and2 == CFALSE) {g0061I_try00622 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          g0061I_try00622 = EID{CTRUE.Id(),0}/* arg-4 */} 
        /* arg-3 */} 
      }
      /* and-2 */} 
    /* ERROR PROTECTION INSERTED (g0061I-Result) */
    if ErrorIn(g0061I_try00622) {Result = g0061I_try00622
    } else {
    g0061I = ToBoolean(OBJ(g0061I_try00622))
    if (g0061I == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *CompileCCast   = To_CompileCCast(new(CompileCCast).Is(C_Compile_C_cast))
        /* noccur = 4 */
        _CL_obj.Arg = x
        _CL_obj.SetArg = s
        Result = EID{_CL_obj.Id(),0}
        /* Let-3 */} 
      } else {
      Result = x.ToEID()
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: Compile/c_strict_check @ any (throw: true) 
func E_Compile_c_strict_check_any (x EID,s EID) EID { 
    return /*(sm for Compile/c_strict_check @ any= EID)*/ F_Compile_c_strict_check_any(ANY(x),ToClass(OBJ(s)) )} 
  
// using conversions. s is a sort or void (we do not need the value).
// note: we need s to be the precise sort for C++
// the is the default version that uses c_code(x)/ c_sort(x)
// in CLAIRE 4, we do not generate conversion at optim time
/* {1} OPT.The go function for: c_code(x:any,s:class) [] */
func F_c_code_any1 (x *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var y *ClaireAny  
      /* noccur = 3 */
      var y_try00803 EID 
      if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0068 *Language.Call   = Language.To_Call(x)
          /* noccur = 1 */
          y_try00803 = F_Optimize_c_code_call_Call(g0068,s)
          /* Let-4 */} 
        } else {
        y_try00803 = Core.F_CALL(C_c_code,ARGS(x.ToEID()))
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (y-Result) */
      if ErrorIn(y_try00803) {Result = y_try00803
      } else {
      y = ANY(y_try00803)
      /* Let:3 */{ 
        var z *ClaireClass  
        /* noccur = 1 */
        var z_try00814 EID 
        z_try00814 = Core.F_CALL(C_Compile_c_sort,ARGS(y.ToEID()))
        /* ERROR PROTECTION INSERTED (z-Result) */
        if ErrorIn(z_try00814) {Result = z_try00814
        } else {
        z = ToClass(OBJ(z_try00814))
        if ((s.Id() == C_void.Id()) || 
            (z.Id() == s.Id())) /* If:4 */{ 
          var g0082I *ClaireBoolean  
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Equal(s.Id(),C_void.Id())
            if (v_and5 == CFALSE) {g0082I = CFALSE
            } else /* arg:6 */{ 
              if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0070 *Language.Call   = Language.To_Call(x)
                  /* noccur = 1 */
                  v_and5 = Equal(g0070.Selector.Id(),C__equal.Id())
                  /* Let-8 */} 
                } else {
                v_and5 = CFALSE
                /* If-7 */} 
              if (v_and5 == CFALSE) {g0082I = CFALSE
              } else /* arg:7 */{ 
                g0082I = CTRUE/* arg-7 */} 
              /* arg-6 */} 
            /* and-5 */} 
          if (g0082I == CTRUE) /* If:5 */{ 
            F_Compile_warn_void()
            Core.F_tformat_string(MakeString("-- Equality meant as an assignment: ~S [264]\n"),2,MakeConstantList(x))
            /* If-5 */} 
          Result = y.ToEID()
          } else {
          Result = y.ToEID()
          /* If-4 */} 
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ list<type_expression>(any, class) (throw: true) 
func E_c_code_any1 (x EID,s EID) EID { 
    return /*(sm for c_code @ list<type_expression>(any, class)= EID)*/ F_c_code_any1(ANY(x),ToClass(OBJ(s)) )} 
  
// basic code generation
// c_code without a sort parameter means that we do not care about the resulting sort,
// which will be checked later on using c_sort
/* {1} OPT.The go function for: c_code(self:any) [] */
func F_c_code_any2 (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0083 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        /* noccur = 1 */
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[215] the symbol ~A is unbound_symbol").Id(),0},EID{g0083.Name.Id(),0}))
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0084 *ClaireVariable   = To_Variable(self)
        /* noccur = 1 */
        Result = EID{g0084.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0085 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
        /* noccur = 2 */
        F_Optimize_c_register_object(ToObject(g0085.Id()))
        Result = EID{g0085.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Optimized_instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0086 *Language.OptimizedInstruction   = Language.To_OptimizedInstruction(self)
        /* noccur = 1 */
        Result = EID{g0086.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0087 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 1 */
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[internal] c_code(~S) should have a parameter").Id(),0},EID{g0087.Id(),0}))
        /* Let-3 */} 
      /* If!2 */}  else if (C_set.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0088 *ClaireSet   = ToSet(self)
        /* noccur = 5 */
        if (F_boolean_I_any(g0088.Id()) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var x *Language.Set  
            /* noccur = 3 */
            /* Let:6 */{ 
              var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
              /* noccur = 3 */
              _CL_obj.Args = g0088.List_I()
              x = _CL_obj
              /* Let-6 */} 
            if (ToList(g0088.Id()).Of().Id() != C_void.Id()) /* If:6 */{ 
              x.Of = ToList(g0088.Id()).Of()
              /* If-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{x.Id(),0}))
            /* Let-5 */} 
          } else {
          Result = EID{g0088.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0089 *ClaireList   = ToList(self)
        /* noccur = 5 */
        if (g0089.Length() != 0) /* If:4 */{ 
          /* Let:5 */{ 
            var x *Language.List  
            /* noccur = 3 */
            /* Let:6 */{ 
              var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
              /* noccur = 3 */
              _CL_obj.Args = g0089
              x = _CL_obj
              /* Let-6 */} 
            if (g0089.Of().Id() != C_void.Id()) /* If:6 */{ 
              x.Of = g0089.Of()
              /* If-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{x.Id(),0}))
            /* Let-5 */} 
          } else {
          Result = EID{g0089.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_tuple.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0090 *ClaireTuple   = ToTuple(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0093UU *Language.Tuple  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Tuple   = Language.To_Tuple(new(Language.Tuple).Is(Language.C_Tuple))
            /* noccur = 3 */
            _CL_obj.Args = g0090.List_I()
            g0093UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g0093UU.Id(),0}))
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      if (self.Isa.IsIn(C_thing) == CTRUE) /* If:3 */{ 
        Core.F_CALL(C_Optimize_c_register,ARGS(self.ToEID()))
        /* If-3 */} 
      Result = self.ToEID()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: c_code @ list<type_expression>(any) (throw: true) 
func E_c_code_any2 (self EID) EID { 
    return /*(sm for c_code @ list<type_expression>(any)= EID)*/ F_c_code_any2(ANY(self) )} 
  
// suggestion for claire4 : get rid of c_sort
/* {1} OPT.The go function for: get_sort(self:any) [] */
func F_Optimize_get_sort_any (self *ClaireAny ) EID { 
    var Result EID 
    Result = Language.F_static_type_any(self)
    return Result} 
  
// The EID go function for: get_sort @ any (throw: true) 
func E_Optimize_get_sort_any (self EID) EID { 
    return /*(sm for get_sort @ any= EID)*/ F_Optimize_get_sort_any(ANY(self) )} 
  
// gives the sort of a compiled expression (does not apply to instructions that
// have a direct c_code(x,s)
// v2.4.9: special type => special sorts !!!
/* {1} OPT.The go function for: Compile/c_sort(self:any) [] */
func F_Compile_c_sort_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0094 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
        /* noccur = 4 */
        if (F_Compile_nativeVar_ask_global_variable(g0094) == CTRUE) /* If:4 */{ 
          if (Equal(g0094.Range.Id(),CEMPTY.Id()) == CTRUE) /* If:5 */{ 
            Result = EID{F_Compile_osort_any(g0094.Value.Isa.Id()).Id(),0}
            } else {
            Result = EID{F_Compile_osort_any(g0094.Range.Id()).Id(),0}
            /* If-5 */} 
          } else {
          Result = EID{C_any.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0095 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 41 */
        if (g0095.Isa.IsIn(C_Variable) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0096 *ClaireVariable   = To_Variable(g0095.Id())
            /* noccur = 1 */
            Result = EID{F_sort_Variable(g0096).Id(),0}
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Assign) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0097 *Language.Assign   = Language.To_Assign(g0095.Id())
            /* noccur = 1 */
            Result = EID{F_sort_Variable(To_Variable(g0097.ClaireVar)).Id(),0}
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Call) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0098 *Language.Call   = Language.To_Call(g0095.Id())
            /* noccur = 1 */
            Result = EID{F_Compile_osort_any(F_Optimize_selector_psort_Call(g0098).Id()).Id(),0}
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0099 *Language.CallMethod   = Language.To_CallMethod(g0095.Id())
            /* noccur = 4 */
            if ((g0099.Arg.Selector.Id() == Core.C_externC.Id()) && 
                (g0099.Args.Length() == 2)) /* If:6 */{ 
              Result = EID{F_Compile_psort_any(g0099.Args.At(2-1)).Id(),0}
              } else {
              Result = F_Optimize_c_srange_method(g0099.Arg)
              /* If-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Call_slot) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0100 *Language.CallSlot   = Language.To_CallSlot(g0095.Id())
            /* noccur = 1 */
            Result = EID{g0100.Selector.Srange.Id(),0}
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Call_table) == CTRUE) /* If:4 */{ 
          Result = EID{C_any.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Call_array) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0102 *Language.CallArray   = Language.To_CallArray(g0095.Id())
            /* noccur = 1 */
            if (g0102.Test == C_float.Id()) /* If:6 */{ 
              Result = EID{C_float.Id(),0}
              } else {
              Result = EID{C_any.Id(),0}
              /* If-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Definition) == CTRUE) /* If:4 */{ 
          Result = EID{C_object.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(C_Compile_C_cast) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0104 *CompileCCast   = To_CompileCCast(g0095.Id())
            /* noccur = 1 */
            Result = EID{g0104.SetArg.Id(),0}
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Update) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0105 *Language.Update   = Language.To_Update(g0095.Id())
            /* noccur = 1 */
            Result = Core.F_CALL(C_Compile_c_sort,ARGS(g0105.Value.ToEID()))
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_If) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0106 *Language.If   = Language.To_If(g0095.Id())
            /* noccur = 2 */
            /* Let:6 */{ 
              var g0127UU *ClaireClass  
              /* noccur = 1 */
              var g0127UU_try01287 EID 
              /* Let:7 */{ 
                var g0129UU *ClaireAny  
                /* noccur = 1 */
                var g0129UU_try01318 EID 
                g0129UU_try01318 = Core.F_CALL(C_Compile_c_sort,ARGS(g0106.Arg.ToEID()))
                /* ERROR PROTECTION INSERTED (g0129UU-g0127UU_try01287) */
                if ErrorIn(g0129UU_try01318) {g0127UU_try01287 = g0129UU_try01318
                } else {
                g0129UU = ANY(g0129UU_try01318)
                /* Let:8 */{ 
                  var g0130UU *ClaireAny  
                  /* noccur = 1 */
                  var g0130UU_try01329 EID 
                  g0130UU_try01329 = Core.F_CALL(C_Compile_c_sort,ARGS(g0106.Other.ToEID()))
                  /* ERROR PROTECTION INSERTED (g0130UU-g0127UU_try01287) */
                  if ErrorIn(g0130UU_try01329) {g0127UU_try01287 = g0130UU_try01329
                  } else {
                  g0130UU = ANY(g0130UU_try01329)
                  g0127UU_try01287 = EID{Core.F_meet_class(ToClass(g0129UU),ToClass(g0130UU)).Id(),0}
                  }
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0127UU-Result) */
              if ErrorIn(g0127UU_try01287) {Result = g0127UU_try01287
              } else {
              g0127UU = ToClass(OBJ(g0127UU_try01287))
              Result = EID{F_Compile_psort_any(g0127UU.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Handle) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0107 *Language.ClaireHandle   = Language.To_ClaireHandle(g0095.Id())
            /* noccur = 2 */
            /* Let:6 */{ 
              var g0133UU *ClaireClass  
              /* noccur = 1 */
              var g0133UU_try01347 EID 
              /* Let:7 */{ 
                var g0135UU *ClaireAny  
                /* noccur = 1 */
                var g0135UU_try01378 EID 
                g0135UU_try01378 = Core.F_CALL(C_Compile_c_sort,ARGS(g0107.Arg.ToEID()))
                /* ERROR PROTECTION INSERTED (g0135UU-g0133UU_try01347) */
                if ErrorIn(g0135UU_try01378) {g0133UU_try01347 = g0135UU_try01378
                } else {
                g0135UU = ANY(g0135UU_try01378)
                /* Let:8 */{ 
                  var g0136UU *ClaireAny  
                  /* noccur = 1 */
                  var g0136UU_try01389 EID 
                  g0136UU_try01389 = Core.F_CALL(C_Compile_c_sort,ARGS(g0107.Other.ToEID()))
                  /* ERROR PROTECTION INSERTED (g0136UU-g0133UU_try01347) */
                  if ErrorIn(g0136UU_try01389) {g0133UU_try01347 = g0136UU_try01389
                  } else {
                  g0136UU = ANY(g0136UU_try01389)
                  g0133UU_try01347 = EID{Core.F_meet_class(ToClass(g0135UU),ToClass(g0136UU)).Id(),0}
                  }
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0133UU-Result) */
              if ErrorIn(g0133UU_try01347) {Result = g0133UU_try01347
              } else {
              g0133UU = ToClass(OBJ(g0133UU_try01347))
              Result = EID{F_Compile_psort_any(g0133UU.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Let) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0108 *Language.Let   = Language.To_Let(g0095.Id())
            /* noccur = 1 */
            Result = Core.F_CALL(C_Compile_c_sort,ARGS(g0108.Arg.ToEID()))
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Return) == CTRUE) /* If:4 */{ 
          Result = EID{C_any.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_List) == CTRUE) /* If:4 */{ 
          Result = EID{C_object.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Set) == CTRUE) /* If:4 */{ 
          Result = EID{C_object.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Tuple) == CTRUE) /* If:4 */{ 
          Result = EID{C_object.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Construct) == CTRUE) /* If:4 */{ 
          Result = EID{C_any.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Gassign) == CTRUE) /* If:4 */{ 
          Result = EID{C_any.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Super) == CTRUE) /* If:4 */{ 
          Result = EID{C_any.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_For) == CTRUE) /* If:4 */{ 
          Result = EID{C_any.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Exists) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0117 *Language.Exists   = Language.To_Exists(g0095.Id())
            /* noccur = 1 */
            if (g0117.Other == CNULL) /* If:6 */{ 
              Result = EID{C_any.Id(),0}
              } else {
              Result = EID{C_object.Id(),0}
              /* If-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Iteration) == CTRUE) /* If:4 */{ 
          Result = EID{C_object.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_And) == CTRUE) /* If:4 */{ 
          Result = EID{C_boolean.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Or) == CTRUE) /* If:4 */{ 
          Result = EID{C_boolean.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_While) == CTRUE) /* If:4 */{ 
          Result = EID{C_any.Id(),0}
          /* If!4 */}  else if (g0095.Isa.IsIn(Language.C_Do) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0122 *Language.Do   = Language.To_Do(g0095.Id())
            /* noccur = 1 */
            /* Let:6 */{ 
              var g0139UU *ClaireAny  
              /* noccur = 1 */
              var g0139UU_try01407 EID 
              g0139UU_try01407 = Core.F_last_list(g0122.Args)
              /* ERROR PROTECTION INSERTED (g0139UU-Result) */
              if ErrorIn(g0139UU_try01407) {Result = g0139UU_try01407
              } else {
              g0139UU = ANY(g0139UU_try01407)
              Result = Core.F_CALL(C_Compile_c_sort,ARGS(g0139UU.ToEID()))
              }
              /* Let-6 */} 
            /* Let-5 */} 
          } else {
          Result = ToException(Core.C_general_error.Make(MakeString("[internal] c_sort is not implemented for ~S").Id(),MakeConstantList(g0095.Id().Isa.Id()).Id())).Close()
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_float.Id() == self.Isa.Id()) /* If:2 */{ 
      Result = EID{C_float.Id(),0}
      } else {
      /* Let:3 */{ 
        var g0141UU *ClaireType  
        /* noccur = 1 */
        var g0141UU_try01424 EID 
        g0141UU_try01424 = Core.F_CALL(C_c_type,ARGS(self.ToEID()))
        /* ERROR PROTECTION INSERTED (g0141UU-Result) */
        if ErrorIn(g0141UU_try01424) {Result = g0141UU_try01424
        } else {
        g0141UU = ToType(OBJ(g0141UU_try01424))
        Result = EID{F_Compile_psort_any(g0141UU.Id()).Id(),0}
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Compile/c_sort @ any (throw: true) 
func E_Compile_c_sort_any (self EID) EID { 
    return /*(sm for Compile/c_sort @ any= EID)*/ F_Compile_c_sort_any(ANY(self) )} 
  
// for the special compiler properties, we need to tell the sort of the optimized
// form
/* {1} OPT.The go function for: selector_psort(self:Call) [] */
func F_Optimize_selector_psort_Call (self *Language.Call ) *ClaireClass  { 
    // procedure body with s =  
var Result *ClaireClass  
    /* Let:2 */{ 
      var p *ClaireProperty   = self.Selector
      /* noccur = 4 */
      if ((p.Id() == Core.C_mClaire_base_I.Id()) || 
          (p.Id() == Core.C_mClaire_index_I.Id())) /* If:3 */{ 
        Result = C_integer
        /* If!3 */}  else if (p.Id() == C_Compile_anyObject_I.Id()) /* If:3 */{ 
        Result = ToClass(self.Args.At(1-1))
        /* If!3 */}  else if (p.Id() == C_Compile_object_I.Id()) /* If:3 */{ 
        Result = ToClass(self.Args.At(2-1))
        } else {
        Result = C_any
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: selector_psort @ Call (throw: false) 
func E_Optimize_selector_psort_Call (self EID) EID { 
    return EID{/*(sm for selector_psort @ Call= class)*/ F_Optimize_selector_psort_Call(Language.To_Call(OBJ(self)) ).Id(),0}} 
  
// ******************************************************************
// *    Part 3: g_throw and status(m:method)                        *
// ******************************************************************
// NEW in claire4 : optimization when compiler.safety is high may prevent throwing exceptions
// these two variabler are used for cross-compiling, when the status changes from the existing(compiled) version to the
// new one being compiled
// NEW in claire 4, because error handling is mananaged by the compiler
// tells if an expression can throw an exception, based on can_throw?(p or m)
/* {1} OPT.The go function for: Compile/g_throw(self:any) [] */
func F_Compile_g_throw_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_bag) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0143 *ClaireBag   = ToBag(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0167UU *ClaireAny  
          /* noccur = 1 */
          var g0167UU_try01685 EID 
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            g0167UU_try01685= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            var x_support_try01696 EID 
            x_support_try01696 = Core.F_enumerate_any(g0143.Id())
            /* ERROR PROTECTION INSERTED (x_support-g0167UU_try01685) */
            if ErrorIn(x_support_try01696) {g0167UU_try01685 = x_support_try01696
            } else {
            x_support = ToList(OBJ(x_support_try01696))
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              var g0194I *ClaireBoolean  
              var g0194I_try01957 EID 
              g0194I_try01957 = F_Compile_g_throw_any(x)
              /* ERROR PROTECTION INSERTED (g0194I-void_try7) */
              if ErrorIn(g0194I_try01957) {void_try7 = g0194I_try01957
              } else {
              g0194I = ToBoolean(OBJ(g0194I_try01957))
              if (g0194I == CTRUE) /* If:7 */{ 
                 /*v = g0167UU_try01685, s =EID*/
g0167UU_try01685 = EID{CTRUE.Id(),0}
                break
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (void_try7-g0167UU_try01685) */
              if ErrorIn(void_try7) {g0167UU_try01685 = void_try7
              g0167UU_try01685 = void_try7
              break
              } else {
              }}
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (g0167UU-Result) */
          if ErrorIn(g0167UU_try01685) {Result = g0167UU_try01685
          } else {
          g0167UU = ANY(g0167UU_try01685)
          Result = EID{F_boolean_I_any(g0167UU).Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Construct) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0144 *Language.Construct   = Language.To_Construct(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0196UU *ClaireAny  
          /* noccur = 1 */
          var g0196UU_try01975 EID 
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            g0196UU_try01975= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = g0144.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              var g0198I *ClaireBoolean  
              var g0198I_try01997 EID 
              g0198I_try01997 = F_Compile_g_throw_any(x)
              /* ERROR PROTECTION INSERTED (g0198I-void_try7) */
              if ErrorIn(g0198I_try01997) {void_try7 = g0198I_try01997
              } else {
              g0198I = ToBoolean(OBJ(g0198I_try01997))
              if (g0198I == CTRUE) /* If:7 */{ 
                 /*v = g0196UU_try01975, s =EID*/
g0196UU_try01975 = EID{CTRUE.Id(),0}
                break
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (void_try7-g0196UU_try01975) */
              if ErrorIn(void_try7) {g0196UU_try01975 = void_try7
              g0196UU_try01975 = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (g0196UU-Result) */
          if ErrorIn(g0196UU_try01975) {Result = g0196UU_try01975
          } else {
          g0196UU = ANY(g0196UU_try01975)
          Result = EID{F_boolean_I_any(g0196UU).Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Assign) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0145 *Language.Assign   = Language.To_Assign(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0145.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Gassign) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0146 *Language.Gassign   = Language.To_Gassign(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0146.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_And) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0147 *Language.And   = Language.To_And(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0147.Args.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Or) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0148 *Language.Or   = Language.To_Or(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0148.Args.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0149 *Language.Call   = Language.To_Call(self)
        /* noccur = 3 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Core.F__I_equal_any(g0149.Selector.Id(),Core.C_unsafe.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try02006 EID 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              var v_or6_try02017 EID 
              v_or6_try02017 = F_Compile_g_throw_any(g0149.Args.Id())
              /* ERROR PROTECTION INSERTED (v_or6-v_and4_try02006) */
              if ErrorIn(v_or6_try02017) {v_and4_try02006 = v_or6_try02017
              } else {
              v_or6 = ToBoolean(OBJ(v_or6_try02017))
              if (v_or6 == CTRUE) {v_and4_try02006 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                var v_or6_try02048 EID 
                v_or6_try02048 = F_Compile_can_throw_ask_property(g0149.Selector)
                /* ERROR PROTECTION INSERTED (v_or6-v_and4_try02006) */
                if ErrorIn(v_or6_try02048) {v_and4_try02006 = v_or6_try02048
                } else {
                v_or6 = ToBoolean(OBJ(v_or6_try02048))
                if (v_or6 == CTRUE) {v_and4_try02006 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  v_and4_try02006 = EID{CFALSE.Id(),0}/* org-8 */} 
                /* org-7 */} 
              }}
              /* or-6 */} 
            /* ERROR PROTECTION INSERTED (v_and4-Result) */
            if ErrorIn(v_and4_try02006) {Result = v_and4_try02006
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try02006))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              Result = EID{CTRUE.Id(),0}/* arg-6 */} 
            /* arg-5 */} 
          }
          /* and-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0150 *Language.CallMethod   = Language.To_CallMethod(self)
        /* noccur = 5 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Core.F__I_equal_any(g0150.Arg.Id(),C_Compile_m_unsafe.Value)
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try02056 EID 
            v_and4_try02056 = F_Optimize_notOpt_Call_method(g0150)
            /* ERROR PROTECTION INSERTED (v_and4-Result) */
            if ErrorIn(v_and4_try02056) {Result = v_and4_try02056
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try02056))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              v_and4 = Core.F__I_equal_any(g0150.Arg.Selector.Id(),Core.C_externC.Id())
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                var v_and4_try02068 EID 
                /* or:8 */{ 
                  var v_or8 *ClaireBoolean  
                  
                  var v_or8_try02079 EID 
                  v_or8_try02079 = F_Compile_g_throw_any(g0150.Args.Id())
                  /* ERROR PROTECTION INSERTED (v_or8-v_and4_try02068) */
                  if ErrorIn(v_or8_try02079) {v_and4_try02068 = v_or8_try02079
                  } else {
                  v_or8 = ToBoolean(OBJ(v_or8_try02079))
                  if (v_or8 == CTRUE) {v_and4_try02068 = EID{CTRUE.Id(),0}
                  } else /* or:9 */{ 
                    var v_or8_try020810 EID 
                    v_or8_try020810 = F_Compile_can_throw_ask_method(g0150.Arg)
                    /* ERROR PROTECTION INSERTED (v_or8-v_and4_try02068) */
                    if ErrorIn(v_or8_try020810) {v_and4_try02068 = v_or8_try020810
                    } else {
                    v_or8 = ToBoolean(OBJ(v_or8_try020810))
                    if (v_or8 == CTRUE) {v_and4_try02068 = EID{CTRUE.Id(),0}
                    } else /* or:10 */{ 
                      v_and4_try02068 = EID{CFALSE.Id(),0}/* org-10 */} 
                    /* org-9 */} 
                  }}
                  /* or-8 */} 
                /* ERROR PROTECTION INSERTED (v_and4-Result) */
                if ErrorIn(v_and4_try02068) {Result = v_and4_try02068
                } else {
                v_and4 = ToBoolean(OBJ(v_and4_try02068))
                if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  Result = EID{CTRUE.Id(),0}/* arg-8 */} 
                /* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          }}
          /* and-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call_slot) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0151 *Language.CallSlot   = Language.To_CallSlot(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0151.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call_table) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0152 *Language.CallTable   = Language.To_CallTable(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0152.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call_array) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0153 *Language.CallArray   = Language.To_CallArray(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0153.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Super) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0154 *Language.Super   = Language.To_Super(self)
        /* noccur = 2 */
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          var v_or4_try02095 EID 
          v_or4_try02095 = F_Compile_g_throw_any(g0154.Args.Id())
          /* ERROR PROTECTION INSERTED (v_or4-Result) */
          if ErrorIn(v_or4_try02095) {Result = v_or4_try02095
          } else {
          v_or4 = ToBoolean(OBJ(v_or4_try02095))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try02106 EID 
            v_or4_try02106 = F_Compile_can_throw_ask_property(g0154.Selector)
            /* ERROR PROTECTION INSERTED (v_or4-Result) */
            if ErrorIn(v_or4_try02106) {Result = v_or4_try02106
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try02106))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              Result = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }}
          /* or-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Update) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0155 *Language.Update   = Language.To_Update(self)
        /* noccur = 2 */
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          var v_or4_try02115 EID 
          v_or4_try02115 = F_Compile_g_throw_any(g0155.Value)
          /* ERROR PROTECTION INSERTED (v_or4-Result) */
          if ErrorIn(v_or4_try02115) {Result = v_or4_try02115
          } else {
          v_or4 = ToBoolean(OBJ(v_or4_try02115))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try02126 EID 
            v_or4_try02126 = F_Compile_g_throw_any(g0155.Arg)
            /* ERROR PROTECTION INSERTED (v_or4-Result) */
            if ErrorIn(v_or4_try02126) {Result = v_or4_try02126
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try02126))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              Result = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }}
          /* or-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Cast) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0156 *Language.Cast   = Language.To_Cast(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0156.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Compile_C_cast) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0157 *CompileCCast   = To_CompileCCast(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0157.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Let) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0158 *Language.Let   = Language.To_Let(self)
        /* noccur = 2 */
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          var v_or4_try02135 EID 
          v_or4_try02135 = F_Compile_g_throw_any(g0158.Value)
          /* ERROR PROTECTION INSERTED (v_or4-Result) */
          if ErrorIn(v_or4_try02135) {Result = v_or4_try02135
          } else {
          v_or4 = ToBoolean(OBJ(v_or4_try02135))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try02146 EID 
            v_or4_try02146 = F_Compile_g_throw_any(g0158.Arg)
            /* ERROR PROTECTION INSERTED (v_or4-Result) */
            if ErrorIn(v_or4_try02146) {Result = v_or4_try02146
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try02146))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              Result = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }}
          /* or-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Do) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0159 *Language.Do   = Language.To_Do(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0159.Args.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_While) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0160 *Language.While   = Language.To_While(self)
        /* noccur = 2 */
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          var v_or4_try02155 EID 
          v_or4_try02155 = F_Compile_g_throw_any(g0160.Test)
          /* ERROR PROTECTION INSERTED (v_or4-Result) */
          if ErrorIn(v_or4_try02155) {Result = v_or4_try02155
          } else {
          v_or4 = ToBoolean(OBJ(v_or4_try02155))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try02166 EID 
            v_or4_try02166 = F_Compile_g_throw_any(g0160.Arg)
            /* ERROR PROTECTION INSERTED (v_or4-Result) */
            if ErrorIn(v_or4_try02166) {Result = v_or4_try02166
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try02166))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              Result = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }}
          /* or-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Construct) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0161 *Language.Construct   = Language.To_Construct(self)
        /* noccur = 1 */
        Result = F_Compile_g_throw_any(g0161.Args.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_If) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0162 *Language.If   = Language.To_If(self)
        /* noccur = 3 */
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          var v_or4_try02175 EID 
          v_or4_try02175 = F_Compile_g_throw_any(g0162.Test)
          /* ERROR PROTECTION INSERTED (v_or4-Result) */
          if ErrorIn(v_or4_try02175) {Result = v_or4_try02175
          } else {
          v_or4 = ToBoolean(OBJ(v_or4_try02175))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try02186 EID 
            v_or4_try02186 = F_Compile_g_throw_any(g0162.Arg)
            /* ERROR PROTECTION INSERTED (v_or4-Result) */
            if ErrorIn(v_or4_try02186) {Result = v_or4_try02186
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try02186))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              var v_or4_try02197 EID 
              v_or4_try02197 = F_Compile_g_throw_any(g0162.Other)
              /* ERROR PROTECTION INSERTED (v_or4-Result) */
              if ErrorIn(v_or4_try02197) {Result = v_or4_try02197
              } else {
              v_or4 = ToBoolean(OBJ(v_or4_try02197))
              if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                Result = EID{CFALSE.Id(),0}/* org-7 */} 
              /* org-6 */} 
            /* org-5 */} 
          }}}
          /* or-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_For) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0163 *Language.For   = Language.To_For(self)
        /* noccur = 2 */
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          var v_or4_try02205 EID 
          v_or4_try02205 = F_Compile_g_throw_any(g0163.SetArg)
          /* ERROR PROTECTION INSERTED (v_or4-Result) */
          if ErrorIn(v_or4_try02205) {Result = v_or4_try02205
          } else {
          v_or4 = ToBoolean(OBJ(v_or4_try02205))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try02216 EID 
            v_or4_try02216 = F_Compile_g_throw_any(g0163.Arg)
            /* ERROR PROTECTION INSERTED (v_or4-Result) */
            if ErrorIn(v_or4_try02216) {Result = v_or4_try02216
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try02216))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              Result = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }}
          /* or-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Iteration) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0164 *Language.Iteration   = Language.To_Iteration(self)
        /* noccur = 2 */
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          var v_or4_try02225 EID 
          v_or4_try02225 = F_Compile_g_throw_any(g0164.SetArg)
          /* ERROR PROTECTION INSERTED (v_or4-Result) */
          if ErrorIn(v_or4_try02225) {Result = v_or4_try02225
          } else {
          v_or4 = ToBoolean(OBJ(v_or4_try02225))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try02236 EID 
            v_or4_try02236 = F_Compile_g_throw_any(g0164.Arg)
            /* ERROR PROTECTION INSERTED (v_or4-Result) */
            if ErrorIn(v_or4_try02236) {Result = v_or4_try02236
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try02236))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              Result = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }}
          /* or-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Handle) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0165 *Language.ClaireHandle   = Language.To_ClaireHandle(self)
        /* noccur = 2 */
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          v_or4 = Core.F__I_equal_any(g0165.Test,C_any.Id())
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try02246 EID 
            v_or4_try02246 = F_Compile_g_throw_any(g0165.Other)
            /* ERROR PROTECTION INSERTED (v_or4-Result) */
            if ErrorIn(v_or4_try02246) {Result = v_or4_try02246
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try02246))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              Result = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }
          /* or-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Compile/g_throw @ any (throw: true) 
func E_Compile_g_throw_any (self EID) EID { 
    return /*(sm for Compile/g_throw @ any= EID)*/ F_Compile_g_throw_any(ANY(self) )} 
  
// return true in regular case, false if the optimization means that no error will occur.
//  this is ugly:  
//    - (x % y) can raise an error in the generic case (using F_belong) but not in the  optimized case
//    - class!(...) can raise an error in interpreted mode nut not at compile time
/* {1} OPT.The go function for: notOpt(self:Call_method) [] */
func F_Optimize_notOpt_Call_method (self *Language.CallMethod ) EID { 
    var Result EID 
    if (self.Arg.Id() == C_Compile_m_member.Value) /* If:2 */{ 
      /* Let:3 */{ 
        var t2 *ClaireClass  
        /* noccur = 4 */
        var t2_try02254 EID 
        t2_try02254 = Language.F_static_type_any(self.Args.At(2-1))
        /* ERROR PROTECTION INSERTED (t2-Result) */
        if ErrorIn(t2_try02254) {Result = t2_try02254
        } else {
        t2 = ToClass(OBJ(t2_try02254))
        Result = EID{MakeBoolean((ToType(t2.Id()).Included(ToType(C_type.Id())) == CTRUE) || (ToType(t2.Id()).Included(ToType(C_list.Id())) == CTRUE) || (ToType(t2.Id()).Included(ToType(C_integer.Id())) == CTRUE) || (ToType(t2.Id()).Included(ToType(C_array.Id())) == CTRUE)).Not.Id(),0}
        }
        /* Let-3 */} 
      /* If!2 */}  else if (self.Arg.Selector.Id() == C_class_I.Id()) /* If:2 */{ 
      Result = EID{self.Args.At(1-1).Isa.IsIn(C_symbol).Not.Id(),0}
      } else {
      Result = EID{CTRUE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: notOpt @ Call_method (throw: true) 
func E_Optimize_notOpt_Call_method (self EID) EID { 
    return /*(sm for notOpt @ Call_method= EID)*/ F_Optimize_notOpt_Call_method(Language.To_CallMethod(OBJ(self)) )} 
  
//regular case !
// can_throw is based on restrictions analysis ... unless it is open => could always return an error
/* {1} OPT.The go function for: Compile/can_throw?(p:property) [] */
func F_Compile_can_throw_ask_property (p *ClaireProperty ) EID { 
    var Result EID 
    /* or:2 */{ 
      var v_or2 *ClaireBoolean  
      
      v_or2 = Equal(MakeInteger(p.Open).Id(),MakeInteger(3).Id())
      if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
      } else /* or:3 */{ 
        var v_or2_try02284 EID 
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Core.F_not_any(ToList(C_Compile_NoErrorOptimize.Value).Memq(p.Id()).Id())
          if (v_and4 == CFALSE) {v_or2_try02284 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try02296 EID 
            /* Let:6 */{ 
              var g0230UU *ClaireAny  
              /* noccur = 1 */
              var g0230UU_try02317 EID 
              /* For:7 */{ 
                var m *ClaireAny  
                _ = m
                g0230UU_try02317= EID{CFALSE.Id(),0}
                for _,m = range(p.Restrictions.ValuesO())/* loop:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  var g0232I *ClaireBoolean  
                  var g0232I_try02339 EID 
                  if (C_method.Id() == m.Isa.Id()) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0226 *ClaireMethod   = ToMethod(m)
                      /* noccur = 1 */
                      g0232I_try02339 = F_Compile_can_throw_ask_method(g0226)
                      /* Let-10 */} 
                    } else {
                    g0232I_try02339 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (g0232I-void_try9) */
                  if ErrorIn(g0232I_try02339) {void_try9 = g0232I_try02339
                  } else {
                  g0232I = ToBoolean(OBJ(g0232I_try02339))
                  if (g0232I == CTRUE) /* If:9 */{ 
                     /*v = g0230UU_try02317, s =EID*/
g0230UU_try02317 = EID{CTRUE.Id(),0}
                    break
                    } else {
                    void_try9 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (void_try9-g0230UU_try02317) */
                  if ErrorIn(void_try9) {g0230UU_try02317 = void_try9
                  g0230UU_try02317 = void_try9
                  break
                  } else {
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (g0230UU-v_and4_try02296) */
              if ErrorIn(g0230UU_try02317) {v_and4_try02296 = g0230UU_try02317
              } else {
              g0230UU = ANY(g0230UU_try02317)
              v_and4_try02296 = EID{F_boolean_I_any(g0230UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_and4-v_or2_try02284) */
            if ErrorIn(v_and4_try02296) {v_or2_try02284 = v_and4_try02296
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try02296))
            if (v_and4 == CFALSE) {v_or2_try02284 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              v_or2_try02284 = EID{CTRUE.Id(),0}/* arg-6 */} 
            /* arg-5 */} 
          }
          /* and-4 */} 
        /* ERROR PROTECTION INSERTED (v_or2-Result) */
        if ErrorIn(v_or2_try02284) {Result = v_or2_try02284
        } else {
        v_or2 = ToBoolean(OBJ(v_or2_try02284))
        if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
        } else /* or:4 */{ 
          Result = EID{CFALSE.Id(),0}/* org-4 */} 
        /* org-3 */} 
      }
      /* or-2 */} 
    return Result} 
  
// The EID go function for: Compile/can_throw? @ property (throw: true) 
func E_Compile_can_throw_ask_property (p EID) EID { 
    return /*(sm for Compile/can_throw? @ property= EID)*/ F_Compile_can_throw_ask_property(ToProperty(OBJ(p)) )} 
  
// access to status ... -1 means that it was never computed 
// Force*Throw is used to adjust for cross-compiling with a status change
/* {1} OPT.The go function for: Compile/can_throw?(m:method) [] */
func F_Compile_can_throw_ask_method (m *ClaireMethod ) EID { 
    var Result EID 
    if (((C_compiler.Safety > 2) && 
          ((ToList(C_Compile_NoErrorOptimize.Value).Memq(m.Id()) == CTRUE) || 
              (ToList(C_Compile_NoErrorOptimize.Value).Memq(m.Selector.Id()) == CTRUE))) || 
        ((m.Isa.IsIn(C_list) == CTRUE) && 
            (ANY(Core.F_CALL(C_of,ARGS(EID{m.Id(),0}))) == C_method.Id()))) /* If:2 */{ 
      Result = EID{CFALSE.Id(),0}
      /* If!2 */}  else if (ToType(C_method.Id()).EmptyList().Memq(m.Id()) == CTRUE) /* If:2 */{ 
      Result = EID{CTRUE.Id(),0}
      /* If!2 */}  else if ((m.Status != -1) || 
        (m.Formula.Id() == CNULL)) /* If:2 */{ 
      Result = EID{Core.F__I_equal_any(MakeInteger(m.Status).Id(),MakeInteger(0).Id()).Id(),0}
      } else {
      Result = F_Compile_can_throw_I_method(m)
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Compile/can_throw? @ method (throw: true) 
func E_Compile_can_throw_ask_method (m EID) EID { 
    return /*(sm for Compile/can_throw? @ method= EID)*/ F_Compile_can_throw_ask_method(ToMethod(OBJ(m)) )} 
  
// here we recursively call g_throw on the body => forced re-compute of status(m) (status!(m) in CLAIRE3)
/* {1} OPT.The go function for: Compile/can_throw!(m:method) [] */
func F_Compile_can_throw_I_method (m *ClaireMethod ) EID { 
    var Result EID 
    m.Status = 0
    var g0234I *ClaireBoolean  
    var g0234I_try02352 EID 
    /* Let:2 */{ 
      var g0236UU *ClaireAny  
      /* noccur = 1 */
      var g0236UU_try02373 EID 
      g0236UU_try02373 = Core.F_CALL(C_c_code,ARGS(m.Formula.Body.ToEID(),EID{m.Range.Class_I().Id(),0}))
      /* ERROR PROTECTION INSERTED (g0236UU-g0234I_try02352) */
      if ErrorIn(g0236UU_try02373) {g0234I_try02352 = g0236UU_try02373
      } else {
      g0236UU = ANY(g0236UU_try02373)
      g0234I_try02352 = F_Compile_g_throw_any(g0236UU)
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (g0234I-Result) */
    if ErrorIn(g0234I_try02352) {Result = g0234I_try02352
    } else {
    g0234I = ToBoolean(OBJ(g0234I_try02352))
    if (g0234I == CTRUE) /* If:2 */{ 
      m.Status = 1
      Result = EID{CTRUE.Id(),0}
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: Compile/can_throw! @ method (throw: true) 
func E_Compile_can_throw_I_method (m EID) EID { 
    return /*(sm for Compile/can_throw! @ method= EID)*/ F_Compile_can_throw_I_method(ToMethod(OBJ(m)) )} 
  
// read can_throw from the status, not influenced by exceptions (for code generation)
// however, for a new method, compute the status
/* {1} OPT.The go function for: Compile/can_throw_status(m:method) [] */
func F_Compile_can_throw_status_method (m *ClaireMethod ) EID { 
    var Result EID 
    if (m.Status == -1) /* If:2 */{ 
      Result = F_Compile_can_throw_I_method(m)
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{Core.F__I_equal_any(MakeInteger(m.Status).Id(),MakeInteger(0).Id()).Id(),0}
    }
    return Result} 
  
// The EID go function for: Compile/can_throw_status @ method (throw: true) 
func E_Compile_can_throw_status_method (m EID) EID { 
    return /*(sm for Compile/can_throw_status @ method= EID)*/ F_Compile_can_throw_status_method(ToMethod(OBJ(m)) )} 
  
// useful #2: provoke a recomputation of status
/* {1} OPT.The go function for: s_throw(m:method) [] */
func F_s_throw_method (m *ClaireMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var la *ClaireLambda   = m.Formula
      /* noccur = 1 */
      /* Let:3 */{ 
        var news *ClaireBoolean  
        /* noccur = 2 */
        var news_try02384 EID 
        news_try02384 = F_Compile_g_throw_any(la.Body)
        /* ERROR PROTECTION INSERTED (news-Result) */
        if ErrorIn(news_try02384) {Result = news_try02384
        } else {
        news = ToBoolean(OBJ(news_try02384))
        Core.F_tformat_string(MakeString("status(~S) := ~S \n"),0,MakeConstantList(m.Id(),news.Id()))
        /* update:4 */{ 
          var va_arg1 *ClaireMethod  
          var va_arg2 int 
          va_arg1 = m
          if (news == CTRUE) /* If:5 */{ 
            va_arg2 = 1
            } else {
            va_arg2 = 0
            /* If-5 */} 
          /* ---------- now we compile update mClaire/status(va_arg1) := va_arg2 ------- */
          va_arg1.Status = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: s_throw @ method (throw: true) 
func E_s_throw_method (m EID) EID { 
    return /*(sm for s_throw @ method= EID)*/ F_s_throw_method(ToMethod(OBJ(m)) )} 
  
// ******************************************************************
// *    Part 4: Names & identifiers management                      *
// ******************************************************************
// check that the module is allowed and otherwise complain because of x;
// this should raise an error, it simply returns false if there is a problem
/* {1} OPT.The go function for: legal?(self:module,x:any) [] */
func F_Optimize_legal_ask_module (self *ClaireModule ,x *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if ((x == C_Compile_object_I.Id()) || 
        (x == C_Compile_anyObject_I.Id())) /* If:2 */{ 
      Result = CTRUE
      /* If!2 */}  else if (F_boolean_I_any(C_OPT.LegalModules.Id()) == CTRUE) /* If:2 */{ 
      var g0241I *ClaireBoolean  
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = C_OPT.LegalModules.Contain_ask(self.Id()).Not
        if (v_and3 == CFALSE) {g0241I = CFALSE
        } else /* arg:4 */{ 
          if (C_method.Id() == x.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0240 *ClaireMethod   = ToMethod(x)
              /* noccur = 1 */
              v_and3 = MakeBoolean((g0240.Inline_ask.Id() == CFALSE.Id()) || (C_compiler.Inline_ask != CTRUE))
              /* Let-6 */} 
            } else {
            v_and3 = CFALSE
            /* If-5 */} 
          if (v_and3 == CFALSE) {g0241I = CFALSE
          } else /* arg:5 */{ 
            g0241I = CTRUE/* arg-5 */} 
          /* arg-4 */} 
        /* and-3 */} 
      if (g0241I == CTRUE) /* If:3 */{ 
        Core.F_tformat_string(MakeString("legal_modules = ~S\n"),0,MakeConstantList(C_OPT.LegalModules.Id()))
        Core.F_tformat_string(MakeString("---- ERROR: ~S implies using ~S !\n\n"),0,MakeConstantList(x,self.Id()))
        Result = CFALSE
        } else {
        Result = CTRUE
        /* If-3 */} 
      } else {
      C_OPT.NeedModules.AddFast(self.Id())
      Result = CTRUE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: legal? @ module (throw: false) 
func E_Optimize_legal_ask_module (self EID,x EID) EID { 
    return EID{/*(sm for legal? @ module= boolean)*/ F_Optimize_legal_ask_module(ToModule(OBJ(self)),ANY(x) ).Id(),0}} 
  
/* {1} OPT.The go function for: legal?(self:environment,x:any) [] */
func F_Optimize_legal_ask_environment (self *ClaireEnvironment ,x *ClaireAny ) *ClaireAny  { 
    // use function body compiling 
return  CTRUE.Id()
    } 
  
// The EID go function for: legal? @ environment (throw: false) 
func E_Optimize_legal_ask_environment (self EID,x EID) EID { 
    return /*(sm for legal? @ environment= any)*/ F_Optimize_legal_ask_environment(ToEnvironment(OBJ(self)),ANY(x) ).ToEID()} 
  
// A named object is used, thus it must be declared if it belongs to the
// current module - returns true if OK
/* {1} OPT.The go function for: c_register(self:(thing U class)) [] */
func F_Optimize_c_register_object (self *ClaireObject ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var x *ClaireAny   = F_Compile_get_module_object(self)
      /* noccur = 2 */
      if (x != ClEnv.Id()) /* If:3 */{ 
        Result = ANY(Core.F_CALL(C_Optimize_legal_ask,ARGS(x.ToEID(),EID{self.Id(),0})))
        } else {
        Result = CTRUE.Id()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_register @ object (throw: false) 
func E_Optimize_c_register_object (self EID) EID { 
    return /*(sm for c_register @ object= any)*/ F_Optimize_c_register_object(ToObject(OBJ(self)) ).ToEID()} 
  
// looks if a property may be implicit and then add it in the right list
/* {1} OPT.The go function for: c_register(self:property) [] */
func F_Optimize_c_register_property (self *ClaireProperty ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var m *ClaireModule   = ClEnv.Module_I
      /* noccur = 1 */
      /* Let:3 */{ 
        var m2 *ClaireAny   = F_Compile_get_module_object(ToObject(self.Id()))
        /* noccur = 2 */
        if (((m2 == C_claire.Id()) || 
              (m2 == m.Id())) && 
            (C_OPT.Objects.Memq(self.Id()) != CTRUE)) /* If:4 */{ 
          C_OPT.Properties.AddFast(self.Id())
          /* If-4 */} 
        Result = F_Optimize_c_register_object(ToObject(self.Id()))
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_register @ property (throw: false) 
func E_Optimize_c_register_property (self EID) EID { 
    return /*(sm for c_register @ property= any)*/ F_Optimize_c_register_property(ToProperty(OBJ(self)) ).ToEID()} 
  
// declare the property as used and check if a property may allocate
/* {1} OPT.The go function for: selector_register(self:property) [] */
func F_Optimize_selector_register_property (self *ClaireProperty ) *ClaireAny  { 
    // use function body compiling 
F_Optimize_c_register_property(self)
    return  self.Id()
    } 
  
// The EID go function for: selector_register @ property (throw: false) 
func E_Optimize_selector_register_property (self EID) EID { 
    return /*(sm for selector_register @ property= any)*/ F_Optimize_selector_register_property(ToProperty(OBJ(self)) ).ToEID()} 
  
// this method looks if the open slot is less than 1 or can be set to 1
// v3.3.48 note - weaken the open semantic to get a better c_status
/* {1} OPT.The go function for: stable?(self:relation) [] */
func F_Optimize_stable_ask_relation (self *ClaireRelation ) *ClaireBoolean  { 
    // use function body compiling 
/* Let:2 */{ 
      var m *ClaireAny   = F_Compile_get_module_object(ToObject(self.Id()))
      /* noccur = 0 */
      _ = m
      if (self.Open == 2) /* If:3 */{ 
        self.Open = 1
        /* If-3 */} 
      /* Let-2 */} 
    return  MakeBoolean((self.Open <= 1) || (self.Open == 4))
    } 
  
// The EID go function for: stable? @ relation (throw: false) 
func E_Optimize_stable_ask_relation (self EID) EID { 
    return EID{/*(sm for stable? @ relation= boolean)*/ F_Optimize_stable_ask_relation(ToRelation(OBJ(self)) ).Id(),0}} 
  
// v3.2.04
// returns the module (i.e. the compilation unit, not the namespace) in which self is
// defined
/* {1} OPT.The go function for: Compile/get_module(self:(thing U class)) [] */
func F_Compile_get_module_object (self *ClaireObject ) *ClaireAny  { 
    // use function body compiling 
return  ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(EID{self.Id(),0})))).Defined().Id()
    } 
  
// The EID go function for: Compile/get_module @ object (throw: false) 
func E_Compile_get_module_object (self EID) EID { 
    return /*(sm for Compile/get_module @ object= any)*/ F_Compile_get_module_object(ToObject(OBJ(self)) ).ToEID()} 
  
//      (while (m.loaded = 0) m := m.part_of, m) ]
// allows to optimize the access
/* {1} OPT.The go function for: known!(l:listargs) [] */
func F_known_I_listargs (l *ClaireList ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    C_OPT.ToRemove.AddFast(Reader.C_known_I.Id())
    /* For:2 */{ 
      var r *ClaireAny  
      _ = r
      Result= CFALSE.Id()
      var r_support *ClaireList  
      r_support = ToList(l.Id())
      r_len := r_support.Length()
      for i_it := 0; i_it < r_len; i_it++ { 
        r = r_support.At(i_it)
        if (r.Isa.IsIn(C_property) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0245 *ClaireProperty   = ToProperty(r)
            /* noccur = 1 */
            C_OPT.Knowns.AddFast(g0245.Id())
            /* Let-5 */} 
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: known! @ listargs (throw: false) 
func E_known_I_listargs (l EID) EID { 
    return /*(sm for known! @ listargs= any)*/ F_known_I_listargs(ToList(OBJ(l)) ).ToEID()} 
  