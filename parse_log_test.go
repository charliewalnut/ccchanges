package main

import (
    "strings"
    "testing"
)

var testLog = []string{`2011-09-18  Dan Bernstein  <mitz@apple.com>

        Leopard build fix after r95391.

        * platform/graphics/mac/ComplexTextControllerATSUI.cpp:
        (WebCore::ComplexTextController::collectComplexTextRunsForCharactersATSUI):

`,`
2011-09-18  Dan Bernstein  <mitz@apple.com>

        Allow Core Text to choose the fallback font for rendering a combining character sequence

        Fixes <rdar://problem/7860281> Decomposed text is displayed incorrectly when Verdana is specified
        https://bugs.webkit.org/show_bug.cgi?id=68287

        Reviewed by Dave Hyatt.

        Test: platform/mac/fast/text/combining-character-sequence-fallback.html

        * WebCore.xcodeproj/project.pbxproj: Compile ComplexTextControllerCoreText.cpp as Objective-C++.
        I am going to rename this file in a followup patch.

        * platform/graphics/Font.h: Declared fontDataForCombiningCharacterSequence().

        * platform/graphics/SimpleFontData.h: Declared canRenderCombiningCharacterSequence() and added
        a member variable to cache the results of calling this function.

        * platform/graphics/mac/ComplexTextController.cpp:
        (WebCore::ComplexTextController::offsetForPosition): Fixed a typo in a comment.
        (WebCore::advanceByCombiningCharacterSequence): Added this helper function, which iterates over
        characters until the end of a sequence of combining characters, zero-width joiners and zero-width
        non-joiners. A single base characater followed by such a (possibly-empty) sequence is known as
        a combining character sequence. This stretches the definition a little because the base character
        may be a combining character as well. Returns false if an invalid UTF-16 sequence is encountered,
        true otherwise.
        (WebCore::ComplexTextController::collectComplexTextRuns): - Replaced the glyphData and nextGlyphData
        variables each with a pair of variables, a SimpleFontData* and a boolean indicating whether a
        character mapped to the missing glyph. This clarifies that this function does not use glyph IDs.
        - Renamed the local variable newC to uppercaseCharacter.
        - Changed to use advanceByCombiningCharacterSequence().
        - Changed to use Font::fontDataForCombiningCharacterSequence() instead of Font::glyphDataForCharacter().
          If there are combining marks and none of the fonts in the fallback list can render the sequence alone,
          use the systemFallbackFontData() constant to indicate to collectComplexTextRunsForCharactersCoreText()
          to allow Core Text to perform its own fallback.
        - Stopped checking isSmallCaps against nextIsSmallCaps. It was redundant, since the small caps FontData
          differs from the normal variant.
        (WebCore::ComplexTextController::collectComplexTextRunsForCharacters): Moved the common logic to
        handle the 0 fontData case from the ATSUI and Core Text continuations of this function into this function.

        * platform/graphics/mac/ComplexTextController.h:
        (WebCore::ComplexTextController::systemFallbackFontData): Added. This constant is used to indicate to
        collectComplexTextRunsForCharactersCoreText() that it should allow Core Text to perform its own font
        fallback.

        * platform/graphics/mac/ComplexTextControllerATSUI.cpp:
        (WebCore::ComplexTextController::collectComplexTextRunsForCharactersATSUI): Handle systemFallbackFontData()
        by using the primary font data.

        * platform/graphics/mac/ComplexTextControllerCoreText.cpp:
        (-[CascadeList initWithFont:WebCore::character:]): Added.
        (-[CascadeList count]): Added.
        (-[CascadeList objectAtIndex:]): Added. Returns an entry from a cascade list of CTFontDescriptorRef
        objects based on the font’s fallback list for the character. The list is initialized lazily.
        (WebCore::ComplexTextController::collectComplexTextRunsForCharactersCoreText): Handle systemFallbackFontData()
        by allowing Core Text to perform font fallback, starting with a cascade list based on the font’s fallback
        list and possibly continuing with system fallback.

        * platform/graphics/mac/FontMac.mm:
        (WebCore::Font::fontDataForCombiningCharacterSequence): Added. If the sequence is only a base character,
        calls through to glyphDataForCharacter(). Otherwise, iterates over the font’s fallback list for the
        base character (and the system fallback font for that character), returning the first font that can
        render the sequence, or 0 if there is none.

        * platform/graphics/mac/SimpleFontDataMac.mm:
        (WebCore::provideStringAndAttributes): Added this Core Text callback.
        (WebCore::SimpleFontData::canRenderCombiningCharacterSequence): Added. Checks if Core Text can render
        the sequence using only this font. Caches the result.

`,`
2011-09-17  Mihai Parparita  <mihaip@chromium.org>

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

`,`
2011-09-17  David Hyatt  <hyatt@apple.com>

        https://bugs.webkit.org/show_bug.cgi?id=68307
        
        Crash in border image cssText. Make sure to null check all the components, since they're all
        optional now.

        Reviewed by Sam Weinig.

        Added fast/borders/border-image-slice-omission.html

        * css/CSSBorderImageValue.cpp:
        (WebCore::CSSBorderImageValue::cssText):

`,`
2011-09-17  Aaron Boodman  <aa@chromium.org>

        Rework script context creation/release notifications
        https://bugs.webkit.org/show_bug.cgi?id=67828

        Reviewed by Adam Barth.

        * bindings/v8/V8DOMWindowShell.cpp:
        (WebCore::V8DOMWindowShell::disposeContextHandles):
        (WebCore::V8DOMWindowShell::initContextIfNeeded):
        * bindings/v8/V8IsolatedContext.cpp:
        (WebCore::V8IsolatedContext::V8IsolatedContext):
        (WebCore::V8IsolatedContext::destroy):
        * bindings/v8/V8IsolatedContext.h:
        * loader/EmptyClients.h:
        (WebCore::EmptyFrameLoaderClient::didCreateScriptContext):
        (WebCore::EmptyFrameLoaderClient::willReleaseScriptContext):
        * loader/FrameLoaderClient.h:

`,`
2011-09-17  Ilya Tikhonovsky  <loislo@chromium.org>

        Web Inspector: file open dialog appears when user clicks on the timeline bar in timeline panel.
        https://bugs.webkit.org/show_bug.cgi?id=68312

        Reviewed by Pavel Feldman.

       * inspector/front-end/TimelinePanel.js:
        (WebInspector.TimelinePanel.prototype._createFileSelector):

`}

func TestParseLog(t *testing.T) {
    entries := parseLog(strings.Join(testLog, "\n"))
    if len(entries) != len(testLog) {
        t.Errorf("expected %d entries but got %d", len(testLog), len(entries))
    }

    for i, e := range entries {
        testEntry(t, parseEntry(strings.Split(testLog[i], "\n")), e)
    }
}
