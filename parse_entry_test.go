package ccchanges

import (
    "testing"
)

func checkPeopleEqual(t *testing.T, actual, expected Person) bool {
    if expected.name != actual.name || expected.email != actual.email {
        t.Errorf("people not equal: got %v expected %v", actual, expected)
        return false
    }
    return true
}

func TestSimpleCommit(t *testing.T) {
    str := `2011-09-17  Mihai Parparita  <mihaip@chromium.org>

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

`
    c := ParseEntry(str)
    mihai := Person{"Mihai Parparita", "mihaip@chromium.org"}
    darin := Person{"Darin Adler", ""}
    if !checkPeopleEqual(t, c.committer, mihai) {
        t.Error("committer not parsed");
    }

    if !checkPeopleEqual(t, c.author, mihai) {
        t.Error("author not parsed");
    }

    if !checkPeopleEqual(t, c.reviewer, darin) {
        t.Error("reviewer not parsed");
    }
}
