var app = new Vue({
  el: '#app',
  data: {
    verses: [
      "*When I star-ted *down _ the _ *street _ last _ *Sun- _ day _, *Fee-lin' migh-ty *low _ and _ *kind- _ a _ *mean. _ ",
      "*Sud-den-ly a *voice _ said _ *\"Go _ forth, _ *neigh- _ bor! _ *Spread the pic-ture *on _ a _ *wi- _ der _ *screen!\" _ And the",
      "*voice _ said, _ *\"Neigh-bor, there's-a *mil- _ lion _ *rea- _ sons _ *why you should be *glad _ in _ *all _ four _ *sea- _ sons! _ ",
      "*Hit the road, _ *neigh-bor leave your *wor-ries and _ *strife! _ *Spread _ the re- *li-gion of the *rhy-thm of _ *life.\" _ For the",
      "*rhy-thm of _ *life _ is-a *pow-er-ful _ *beat, _ Puts-a *tin-gle in your *fin-gers and-a *tin-gle in your *feet! _ ",
      "*Rhy-thm on thi *in- _ side, _ *rhy-thm on the *street, _ and the *rhy-thm of _ *life _ is-a *pow-er-ful _ *beat! _ For the"
    ],
    rows: [],
    countOfPage: 8,
    currPage: 1,
    filter_name: ''
  },
  computed: {
  filteredRows: function(){
    // 因為 JavaScript 的 filter 有分大小寫，
    // 所以這裡將 filter_name 與 rows[n].name 通通轉小寫方便比對。
    var filter_name = this.filter_name.toLowerCase();

    // 如果 filter_name 有內容，回傳過濾後的資料，否則將原本的 rows 回傳。
    return ( this.filter_name.trim() !== '' ) ? 
      this.rows.filter(function(d){ return d.name.toLowerCase().indexOf(filter_name) > -1; }) : 
    this.rows;
  },
    pageStart: function(){
        return (this.currPage - 1) * this.countOfPage;
      },
    totalPage: function(){
      return Math.ceil(this.filteredRows.length / this.countOfPage);
    },
    Verse: function(){
      return this.verses[this.currPage-1].split(" *"); //(this.currPage);
    }

  },
  methods: {
    setPage: function(idx){
      if( idx <= 0 || idx > this.totalPage ){
        return;
      }
      this.currPage = idx;
    },
  },
  created: function(){
    var self = this;
    $.get('http://www.json-generator.com/api/json/get/cknklDscqG?indent=2', function(data){
//      self.rows = data;
    });
  }
});
