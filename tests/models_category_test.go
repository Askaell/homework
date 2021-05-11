package tests

import (
	"testing"

	"github.com/Askaell/homework/pkg/models"
)

//test data
var (
	category1 = models.NewCategory(1, "slug 1", "category 1", nil)
	category2 = models.NewCategory(2, "slug 2", "child of category 1", category1)
	category3 = models.NewCategory(3, "slug 3", "child of category 2", category2)
	category4 = models.NewCategory(4, "slug 4", "child of category 2", category2)
)

func Test_getRootParent1(t *testing.T) {
	if category1.GetRootParent() != category1 {
		t.Error("expected result: ", category1, "\nreal result: ", category1.GetRootParent())
	}
}

func Test_getRootParent2(t *testing.T) {
	if category2.GetRootParent() != category1 {
		t.Error("expected result: ", category1, "\nreal result: ", category2.GetRootParent())
	}
}

func Test_getRootParent3(t *testing.T) {
	if category3.GetRootParent() != category1 {
		t.Error("expected result: ", category1, "\nreal result: ", category3.GetRootParent())
	}
}

func Test_getRootParent4(t *testing.T) {
	if category4.GetRootParent() != category1 {
		t.Error("expected result: ", category1, "\nreal result: ", category4.GetRootParent())
	}
}

func Test_getBreadCrumbs(t *testing.T) {
	expectedResult := models.Categories{*category1, *category2, *category4}
	realResult := category4.GetBreadCrumbs()

	if len(expectedResult) != len(realResult) {
		t.Error("\nincorrect len of result slice\nexpected result: ", len(expectedResult), "\nreal result:", len(realResult))
	}

	for i, category := range realResult {
		if category.Id != expectedResult[i].Id {
			t.Error("expected result: ", expectedResult, "\n real result:", realResult)
		}
	}
}

func Test_getBreadCrumbsNames(t *testing.T) {
	expectedResult := []string{category1.Name, category2.Name, category4.Name}
	realResult := category4.GetBreadCrumbs()

	if len(expectedResult) != len(realResult) {
		t.Error("\nincorrect len of result slice\nexpected result: ", len(expectedResult), "\nreal result:", len(realResult))
	}

	for i, name := range realResult.Names() {
		if name != expectedResult[i] {
			t.Error("expected result: ", expectedResult, "\n real result:", realResult.Names())
		}
	}
}

func Test_getBreadCrumbsSlugs(t *testing.T) {
	expectedResult := []string{category1.Slug, category2.Slug, category4.Slug}
	realResult := category4.GetBreadCrumbs()

	if len(expectedResult) != len(realResult) {
		t.Error("\nincorrect len of result slice\nexpected result: ", len(expectedResult), "\nreal result:", len(realResult))
	}

	for i, slug := range realResult.Slugs() {
		if slug != expectedResult[i] {
			t.Error("expected result: ", expectedResult, "\n real result:", realResult.Slugs())
		}
	}
}
