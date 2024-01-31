package parser

import (
	"errors"

	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/token"
)

func makeLeaf(t token.Token, typ string) ast.ASTree {
	if typ == "" {
		// other token (\n ; def...)
		return ast.NewASTLeaf(t)
	}
	switch typ {
	case "string_literal":
		return ast.NewStringLiteral(t)
	case "number_literal":
		return ast.NewNumberLiteral(t)
	case "name":
		return ast.NewName(t)
	default:
		log.Error("make error leaf")
		return nil
	}
}

func makeAstList(typ string, list *base.List[ast.ASTree]) ast.ASTree {
	if typ == "" {
		if list.Length() == 1 {
			item, err := list.Get(0)
			if err != nil {
				log.Error("make ast list error")
				return nil
			}
			return item
		} else {
			return ast.NewASTList(list)
		}
	}
	switch typ {
	case "include_stmt":
		newList, err := useSecond(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewIncludeStmt(newList)
	case "primary_expr":
		return ast.NewPrimaryExpr(list)
	case "binary_expr":
		return ast.NewBinaryExpr(list)
	case "negative_expr":
		newList, err := useSecond(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewNegativeExpr(newList)
	case "not_expr":
		newList, err := useSecond(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewNotExpr(newList)
	case "null_stmt":
		return ast.NewNullStmt(list)
	case "go_stmt":
		return ast.NewGoStmt(list)
	case "arguments":
		newList, err := ignoreFirst(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewArguments(newList)
	case "block_stmt":
		newList, err := makeBlockStmt(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewBlockStmt(newList)
	case "if_stmt":
		newList, err := parseIfStmt(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewIfStmt(newList)
	case "for_stmt":
		return ast.NewForStmt(list)
	case "while_stmt":
		newList, err := parseWhileStmt(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewWhileStmt(newList)
	case "array_ref":
		newList, err := parseArrayRefStmt(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewArrayRef(newList)
	case "def_stmt":
		return ast.NewDefStmt(list)
	case "return_stmt":
		newList, err := useSecond(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewReturnStmt(newList)
	case "parameter_list":
		newList, err := ignoreFirst(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewParameterList(newList)
	case "array_literal":
		newList, err := ignoreFirst(list)
		if err != nil {
			log.Error(err)
			return nil
		}
		return ast.NewArrayLiteral(newList)
	case "continue_stmt":
		return ast.NewContinueStmt(list)
	case "break_stmt":
		return ast.NewBreakStmt(list)
	default:
		log.Errorf("make error ast list: %s", typ)
		return nil
	}
}

func parseArrayRefStmt(list *base.List[ast.ASTree]) (
	*base.List[ast.ASTree], error) {
	newList := base.NewList[ast.ASTree]()
	if list.Length() < 3 {
		return nil, errors.New("invalid array ref")
	}
	item, err := list.Get(1)
	if err != nil {
		return nil, errors.New("invalid array ref")
	}
	newList.Add(item)
	return newList, nil
}

func makeBlockStmt(list *base.List[ast.ASTree]) (
	*base.List[ast.ASTree], error) {
	newList := base.NewList[ast.ASTree]()
	if list.Length() < 2 {
		return nil, errors.New("invalid block stmt")
	}
	for _, v := range list.Items()[1 : list.Length()-1] {
		vList, ok := v.(*ast.ASTList)
		if !ok {
			return nil, errors.New("invalid ast list")
		}
		if vList.NumChildren() != 2 {
			return nil, errors.New("invalid children")
		}
		val, err := vList.Child(1)
		if err != nil {
			return nil, err
		}
		newList.Add(val)
	}
	return newList, nil
}

func useSecond(list *base.List[ast.ASTree]) (
	*base.List[ast.ASTree], error) {
	newList := base.NewList[ast.ASTree]()
	if list.Length() < 2 {
		return nil, errors.New("invalid ast length")
	}
	first, err := list.Get(1)
	if err != nil {
		return nil, err
	}
	newList.Add(first)
	return newList, nil
}

func parseIfStmt(list *base.List[ast.ASTree]) (
	*base.List[ast.ASTree], error) {
	newList := base.NewList[ast.ASTree]()
	if list.Length() == 1 {
		first, err := list.Get(0)
		if err != nil {
			return nil, err
		}
		newList.Add(first)
	} else if list.Length() == 3 {
		for _, v := range list.Items()[1:] {
			vList, ok := v.(ast.ASTree)
			if !ok {
				return nil, errors.New("invalid ast list")
			}
			newList.Add(vList)
		}
	} else if list.Length() == 4 {
		for _, v := range list.Items()[1:3] {
			vList, ok := v.(ast.ASTree)
			if !ok {
				return nil, errors.New("invalid ast list")
			}
			newList.Add(vList)
		}
		elseT, err := list.Get(3)
		if err != nil {
			return nil, errors.New("invalid if stmt")
		}
		elseBlockList, isBlock := elseT.(*ast.ASTList)
		if !isBlock {
			return nil, errors.New("invalid if stmt")
		}
		if elseBlockList.NumChildren() != 2 {
			return nil, errors.New("invalid else stmt")
		}
		elseBlock, err := elseBlockList.Child(1)
		if err != nil {
			return nil, errors.New("invalid else stmt")
		}
		newList.Add(elseBlock)
	} else {
		return nil, errors.New("invalid parameter list")
	}
	return newList, nil
}

func parseWhileStmt(list *base.List[ast.ASTree]) (
	*base.List[ast.ASTree], error) {
	newList := base.NewList[ast.ASTree]()
	if list.Length() == 1 {
		first, err := list.Get(0)
		if err != nil {
			return nil, err
		}
		newList.Add(first)
	} else if list.Length() == 3 {
		for _, v := range list.Items()[1:] {
			vList, ok := v.(ast.ASTree)
			if !ok {
				return nil, errors.New("invalid ast list")
			}
			newList.Add(vList)
		}
	} else {
		return nil, errors.New("invalid parameter list")
	}
	return newList, nil
}

func ignoreFirst(list *base.List[ast.ASTree]) (
	*base.List[ast.ASTree], error) {
	newList := base.NewList[ast.ASTree]()
	if list.Length() == 0 {
		return newList, nil
	}
	if list.Length() == 1 {
		first, err := list.Get(0)
		if err != nil {
			return nil, err
		}
		newList.Add(first)
	} else if list.Length() > 1 {
		first, err := list.Get(0)
		if err != nil {
			return nil, err
		}
		newList.Add(first)
		for _, v := range list.Items()[1:] {
			vList, ok := v.(*ast.ASTList)
			if !ok {
				return nil, errors.New("invalid ast list")
			}
			if vList.NumChildren() != 2 {
				return nil, errors.New("invalid children")
			}
			val, err := vList.Child(1)
			if err != nil {
				return nil, err
			}
			newList.Add(val)
		}
	} else {
		return nil, errors.New("invalid parameter list")
	}
	return newList, nil
}
