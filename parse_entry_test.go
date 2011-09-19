package main

import (
    "testing"
    "strings"
)

func checkEqual(t *testing.T, expected, actual, field string) {
    if expected != actual {
        t.Errorf("expected %s but got %s for %s", expected, actual, field)
    }
}

func comparePaths(expected, actual []string) bool {
    if len(actual) != len(expected) {
        return false
    }
    for i, a := range actual {
        if expected[i] != a {
            return false
        }
    }
    return true
}

func parseAndTestEntry(t *testing.T, expected Change, entry string) {
    testEntry(t, expected, parseEntry(strings.Split(entry, "\n")))
}

func testEntry(t *testing.T, expected, actual Change) {
    checkEqual(t, expected.committer, actual.committer, "committer");
    checkEqual(t, expected.author, actual.author, "author")
    checkEqual(t, expected.reviewer, actual.reviewer, "reviewer")
    if !comparePaths(expected.paths, actual.paths) {
        t.Errorf("expected %v but got %v for paths", expected.paths, actual.paths)
    }
    if actual.rollout != expected.rollout {
        t.Errorf("expected %v but got %v for rollout", expected.rollout, actual.rollout)
    }
}

// Some handy names for test cases
const mihai = "Mihai Parparita"
const darin = "Darin Adler"
const mitz = "Dan Bernstein"
const rniwa = "Ryosuke Niwa"
const sheriffBot = "Sheriff Bot"

func TestSimple(t *testing.T) {
    paths := []string{"history/BackForwardListImpl.cpp", "loader/EmptyClients.h", "loader/FrameLoaderClient.h"}
    expected := Change{committer: mihai, author: mihai, reviewer: darin, paths: paths}
    parseAndTestEntry(t, expected, `2011-09-17  Mihai Parparita  <mihaip@chromium.org>

        FrameLoaderClient BackForwardList-related methods are unsued
        https://bugs.webkit.org/show_bug.cgi?id=68293

        Reviewed by Darin Adler.

        Remove FrameLoaderClient methods that were added by r51629, since only
        the old (since-deleted) Android port needed them.

        * history/BackForwardListImpl.cpp:
        (WebCore::BackForwardListImpl::addItem):
        (WebCore::BackForwardListImpl::goBack):
        (WebCore::BackForwardListImpl::goForward):
        (WebCore::BackForwardListImpl::goToItem):
        (WebCore::BackForwardListImpl::setCapacity):
        * loader/EmptyClients.h:
        * loader/FrameLoaderClient.h:

`)
}

func TestUnreviewed(t *testing.T) {
    paths := []string{"WebCore.gyp/WebCore.gyp"}
    expected := Change{committer: mitz, author: mitz, paths: paths}
    parseAndTestEntry(t, expected, `2011-09-18  Dan Bernstein  <mitz@apple.com>

        Try to fix the Chromium Mac build after r95391.

        * WebCore.gyp/WebCore.gyp:

`)
}

func TestRollout(t *testing.T) {
    paths := []string{"bindings/scripts/CodeGeneratorGObject.pm"}
    expected := Change{committer: sheriffBot, author: sheriffBot, rollout: true, paths: paths}
    parseAndTestEntry(t, expected, `2011-09-16  Sheriff Bot  <webkit.review.bot@gmail.com>

        Unreviewed, rolling out r95304.
        http://trac.webkit.org/changeset/95304
        https://bugs.webkit.org/show_bug.cgi?id=68299

        Broke GTK+ builds. It pulls derived headers into main.c
        (Requested by rniwa on #webkit).

        * bindings/scripts/CodeGeneratorGObject.pm:

`)
}

