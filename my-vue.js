var app = new Vue({
  el: '#app',
  data: {
    verses: [ "",
      "When I star-ted *down _ the _ *street _ last _ *Sun- _ day _, *Fee-lin' migh-ty *low _ and _ *kind- _ a _ *mean. _ ",
      "Sud-den-ly a *voice _ said _ *\"Go _ forth, _ *neigh- _ bor! _ *Spread the pic-ture *on _ a _ *wi- _ der _ *screen!\" _ And the",
      "voice _ said, _ *\"Neigh-bor, there's-a *mil- _ lion _ *rea- _ sons _ *why you should be *glad _ in _ *all _ four _ *sea- _ sons! _ ",
      "Hit the road, _ *neigh-bor leave your *wor-ries and _ *strife! _ *Spread _ the re- *li-gion of the *rhy-thm of _ *life.\" _ For the",
      "rhy-thm of _ *life _ is-a *pow-er-ful _ *beat, _ Pu-tsa *tin-gle in your *fin-gers and-a *tin-gle in your *feet! _ ",
      "Rhy-thm on thi *in- _ side, _ *rhy-thm on the *street, _ and the *rhy-thm of _ *life _ is-a *pow-er-ful _ *beat! _ For the"
    ],
    currBar: 1,
    countOfBar: 8,
    totalPage: 6,
    currPage: 1,
    bpm: 72
  },
  computed: {
    Verse: function(){
      return this.verses[this.currPage].split("*");
    }
  },
  methods: {
    setBar: function(idx){
      if( idx <= 0 || idx > this.countOfBar ){
        return;
      }
      this.currBar = idx;
    },
    setPage: function(idx){
      if( idx <= 0 || idx > this.totalPage ){
        return;
      }
      this.currPage = idx;
      this.currBar = 1;
    }
  },
  directives: {
    numberOnly: {
      bind: function(el) {
        el.handler = function() {
          el.value = el.value.replace(/\D+/, '')
        }
        el.addEventListener('input', el.handler)
      },
      unbind: function(el) {
        el.removeEventListener('input', el.handler)
      }
    }
  }
});
