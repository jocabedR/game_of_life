<template>
  <h1>Game of life</h1>

  <button @click="setDimensions(10)">10x10</button>
  <button @click="setDimensions(20)">20x20</button>
  <button @click="setDimensions(50)">50x50</button>
  <button @click="setDimensions(100)">100x100</button>

  <form class="options" @submit.prevent="setGrid">
    <br>Probability: <input type="number" min="0" max="1" step="0.1" v-model="probability"/>
    <br><button>Generate a seed!</button>
  </form>

  <button @click="playOrStop">{{play_stop}}</button>
  
  <Grid :grid="grid"/>

</template>

<script>
import axios from 'axios'
import Grid from './components/Grid.vue'

var n = 0
var state

export default {
  name: 'App',
  components: {
    Grid
  },

  data() {
    return{
      size: 0,
      probability: 0.0,
      grid: [],
      play_stop: "Play"
    }
  },

  methods:{
    setDimensions(size){
      this.size = size
    },

    setGrid() {
      if (this.size == 0 || this.probability == 0.0 ) {
        alert('Please chose any size and/or set the probabiliy.')
      } else {
        for(let i=0; i<this.size; i++){
          this.grid[i] = []
          for (let j = 0; j < this.size; j++) {
            if(this.randomNumber() > (this.probability)*10){
              this.grid[i][j] = false 
            } 
            else{
              this.grid[i][j] = true
            } 
          }
        }
      }
    },

    randomNumber() {
      return Math.floor(Math.random() * (10 - 1 + 1)) + 1
    },

    postreq: function() {
      var data = {"grid": this.grid}
      axios({ method: "POST", url: "http://127.0.0.1:3000/game", data: data, headers: {"content-type": "text/plain" } }).then(result => { 
          this.grid = result.data
        }).catch( error => {
            console.error(error);
      });
      
    },

    playOrStop() {
      if (this.size == []) {
        alert('Please chose generate a seed.')
      } else {
        if(this.play_stop == "Play!"){
          this.play_stop = "Stop"
          state = setInterval(this.postreq, 100);
        } else {
          this.play_stop = "Play!"
          clearInterval(state)
        }
      }
    },

  }
}
</script>

<style>
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 3rem;
  }

  .number-input, button{
    margin: 0.5rem 0.25rem;
  }
</style>