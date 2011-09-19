package ccchanges

import (
    "testing"
)

func checkEqual(t *testing.T, expected, actual, field string) {
    if expected != actual {
        t.Errorf("expected %s but got %s for %s", expected, actual, field)
    }
}

func parseAndTestEntry(t *testing.T, expected Change, entry string) {
    c := ParseEntry(entry)

    checkEqual(t, c.committer, expected.committer, "committer");
    checkEqual(t, c.author, expected.author, "author")
    checkEqual(t, c.reviewer, expected.reviewer, "reviewer")
    if c.rollout != expected.rollout {
        t.Errorf("expected %v but got %v for rollout", c.rollout, expected.rollout)
    }
}

// Some handy names for test cases
const mihai = "Mihai Parparita"
const darin = "Darin Adler"
const mitz = "Dan Bernstein"
const rniwa = "Ryosuke Niwa"
const sheriffBot = "Sheriff Bot"

func TestSimple(t *testing.T) {
    expected := Change{committer: mihai, author: mihai, reviewer: darin}
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
    expected := Change{committer: mitz, author: mitz}
    parseAndTestEntry(t, expected, `2011-09-18  Dan Bernstein  <mitz@apple.com>

        Try to fix the Chromium Mac build after r95391.

        * WebCore.gyp/WebCore.gyp:

`)
}

func TestRollout(t *testing.T) {
    expected := Change{committer: sheriffBot, author: sheriffBot, rollout: true}
    parseAndTestEntry(t, expected, `2011-09-16  Sheriff Bot  <webkit.review.bot@gmail.com>

        Unreviewed, rolling out r95304.
        http://trac.webkit.org/changeset/95304
        https://bugs.webkit.org/show_bug.cgi?id=68299

        Broke GTK+ builds. It pulls derived headers into main.c
        (Requested by rniwa on #webkit).

        * bindings/scripts/CodeGeneratorGObject.pm:

`)
}

