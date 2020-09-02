package main

import (
	"fmt"
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func main() {
	converter := md.NewConverter("", true, nil)

	html := `<div class="grid--item body body__container article__body grid-layout__content"><p class="has-dropcap has-dropcap__lead-standard-heading">On Friday, while prosecutors working for the special counsel, Robert Mueller, obtained their first grand-jury indictments in their investigation of potential collusion by the Trump campaign and Russia, the President of the United States was busy gaslighting. Trump tweeted, of course, that “It is now commonly agreed, after many months of <em class="small">COSTLY</em> looking, that there was <em class="small">NO</em> collusion between Russia and Trump. Was collusion with HC!”</p><p>The President was referring to an episode that took place in 2010 whereby the Obama Administration gave a Russian firm permission to buy a Canadian company that had the rights to mine a great deal of uranium in the U.S. The allegation is that, because Bill Clinton took some money from Russian interests, Hillary Clinton, in exchange, approved the uranium deal. Assume for a moment that Russian influence did affect the Obama Administration’s decision. This is not out of the realm of possibility, and the Obama Justice Department was reportedly looking into Russian influence in the uranium markets.</p><div class="consumer-marketing-unit consumer-marketing-unit--article-mid-content consumer-marketing-unit--no-failsafe" role="presentation" aria-hidden="true"><div class="consumer-marketing-unit__slot consumer-marketing-unit__slot--article-mid-content consumer-marketing-unit__slot--in-content"></div><div class="journey-unit"></div></div><p>But the vast majority of examinations by journalists of the uranium deal have found no sign of wrongdoing. I won’t rehash all the details of the bureaucratic process that led to the deal being approved, but it required the support of multiple government agencies, including the Nuclear Regulatory Commission, a body not controlled by Hillary Clinton. If you want the full explanation for why this allegation is false, I highly recommend this detailed <a class="external-link" data-event-click="{&quot;element&quot;:&quot;ExternalLink&quot;,&quot;outgoingURL&quot;:&quot;http://www.factcheck.org/2017/10/facts-uranium-one/&quot;}" href="http://www.factcheck.org/2017/10/facts-uranium-one/" rel="nofollow noopener" target="_blank">account</a> from FactCheck.org, which concludes, “Donald Trump falsely accused former Secretary of State Hillary Clinton of giving away U.S. uranium rights to the Russians and claimed—without evidence—that it was done in exchange for donations to the Clinton Foundation.”</p><div class="ad ad--in-content"><div class="ad__slot ad__slot--in-content"></div></div><p>What’s more important is that Trump is once again spreading lies to confuse the public about the Russian attack on American democracy last year. There are some obvious reasons why Trump would make this untruthful claim. The first is political. Trump’s typical response to any allegation of wrongdoing is to accuse his accuser of the same crime. Perhaps the most famous moment of the Presidential debates last year was Trump’s response when Hillary Clinton accused him of being Vladimir Putin’s puppet. “No puppet, no puppet, you’re the puppet,” he muttered into his microphone. He has been trying to make that case ever since.</p><p>The time line of Hillary Clinton’s relationship with Putin after the uranium deal is hard to square with Trump’s accusation. Let’s start with the top-line conclusion of the Intelligence Community’s assessment about Russian interference in the 2016 election, an assessment that Trump and the top officials of his Administration have—however reluctantly—on occasion publicly endorsed:</p><blockquote class="blockquote-embed blockquote-embed--has-paragraph-margin blockquote-embed--has-small-margins"><div class="blockquote-embed__content"><p>Russian President Vladimir Putin ordered an influence campaign in 2016
aimed at the US presidential election, the consistent goals of which
were to undermine public faith in the U.S. democratic process,
denigrate Secretary Clinton, and harm her electability and potential
presidency. We further assess Putin and the Russian Government
developed a clear preference for President-elect Trump. When it
appeared to Moscow that Secretary Clinton was likely to win the
election, the Russian influence campaign then focused on undermining
her expected presidency.</p></div></blockquote><p>The roots of Putin’s animosity toward Clinton and preference for Trump are clear. Putin, according to the intelligence community, <a class="external-link" data-event-click="{&quot;element&quot;:&quot;ExternalLink&quot;,&quot;outgoingURL&quot;:&quot;http://www.nytimes.com/2011/12/09/world/europe/putin-accuses-clinton-of-instigating-russian-protests.html&quot;}" href="http://www.nytimes.com/2011/12/09/world/europe/putin-accuses-clinton-of-instigating-russian-protests.html" rel="nofollow noopener" target="_blank">blamed Clinton</a> for stoking protests against his regime in 2011 after Russian parliamentary elections that were called fraudulent by independent observers. Coming in the wake of the Arab Spring, the protests were more widespread than anything Putin had previously faced. The Russian leader publicly accused Clinton of being behind them, charging at the time, as the <em>Times</em> reported, that she sent “a signal” to “some actors in our country.” The U.S. intelligence report concluded, “Putin most likely wanted to discredit Secretary Clinton because he has publicly blamed her since 2011 for inciting mass protests against his regime in late 2011 and early 2012, and because he holds a grudge for comments he almost certainly saw as disparaging him.”</p><div class="cne-interlude-embed"></div><p>Putin also believed that the Panama Papers, which disclosed a global offshore accounting network and implicated close friends of Putin, was secretly directed by the U.S. government. The Russian investigative journalists Irina Borogan and Andrei Soldatov believe Putin and his national-security advisers decided on April 8, 2016, to retaliate against the Obama Administration for what Putin believed was America’s role in the Panama Papers disclosures. “It was seen as an attack on personal friends of Putin, his immediate circle,” the authors told the Washington <em>Post</em>. “It’s a line you cannot cross with Putin.” (One of the reporters who led the investigation—Daphne Caruana Galizia—was recently murdered in a car bombing in Malta.) The Russian authors also told the <em>Post</em>, “Putin believed the Panama Papers attack was sponsored by Hillary Clinton’s people.”</p><p>There is no ambiguity about Russia’s preference in the election, and the only reason it needs to be reiterated is that Trump regularly lies about this basic fact. “We assess the influence campaign aspired to help President-elect Trump’s chances of victory when possible by discrediting Secretary Clinton and publicly contrasting her unfavorably to the President-elect,” the joint report by sixteen U.S. intelligence agencies concluded.</p><p>The second, related, reason for Trump to make this false allegation is that he needs some Democratic “scandal” for his supporters on the Hill and in the media to feast on as the details of the Mueller probe become more problematic for his Administration. And what is more alarming than Trump’s lies about this issue is his ability to get top Republicans and influential conservative media institutions to rally around the idea that the Clinton uranium deal is the real collusion story Americans should care about. Congressional Republicans have announced hearings on the matter and Fox News is devoting a large quantity of its coverage to the issue. The frightening aspect of this is that Trump is able to confuse enough supporters into seeing the uranium episode as more important than the documented history of a massive Russian interference campaign in the election and the still unanswered questions about potential collusion between his campaign and Russian actors. It’s yet another reason that a comprehensive, bipartisan accounting by the Republican-controlled Congress of what happened in 2016 is slipping out of reach. Of course, Trump’s ability to gaslight the public has its limits. The indictments obtained by Mueller’s staffers are expected to be unsealed—or leak—by Monday. They are not expected to concern the Clintons or uranium.</p></div><div class="grid--item split-ad-rail grid-layout__aside"><aside class="persistent-aside" style="position:absolute;top:auto;height:auto"><div class="sticky-box"><div class="sticky-box__primary"><div class="split-ad-rail-content"><div class="split-ad-rail-top"><div class="sticky-box"><div class="sticky-box__primary"><div class="ad ad--rail"><div class="ad__slot ad__slot--rail"></div></div><div class="consumer-marketing-unit consumer-marketing-unit--display-rail consumer-marketing-unit--no-failsafe" role="presentation" aria-hidden="true"><div class="consumer-marketing-unit__slot consumer-marketing-unit__slot--display-rail"></div><div class="journey-unit"></div></div></div><div class="sticky-box__placeholder"></div></div></div><div class="split-ad-rail-middle"></div><div class="split-ad-rail-bottom"><div class="sticky-box"><div class="sticky-box__primary"><div class="ad ad--rail"><div class="ad__slot ad__slot--rail"></div></div><div class="consumer-marketing-unit consumer-marketing-unit--display-rail consumer-marketing-unit--no-failsafe" role="presentation" aria-hidden="true"><div class="consumer-marketing-unit__slot consumer-marketing-unit__slot--display-rail"></div><div class="journey-unit"></div></div></div><div class="sticky-box__placeholder"></div></div></div></div></div><div class="sticky-box__placeholder"></div></div></aside></div>`

	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("md ->", markdown)
}
