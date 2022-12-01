<template>
  <h1>Game of life</h1>

  <button @click="setDimensions(10)">10x10</button>
  <button @click="setDimensions(20)">20x20</button>
  <button @click="setDimensions(50)">50x50</button>
  <button @click="setDimensions(100)">100x100</button>

  <form class="options" @submit.prevent="setGrid">
    <br>Probability: <input type="number" min="0" max="1" step="0.1" v-model="probability"/>
    <br><button>Play!</button>
  </form>
  
  <table>
    <tr v-for="(row, rowKey) in cells" :key="rowKey">
      <td v-for="(col, colKey) in row" :key="colKey" 
      :class="col ? 'alive' : 'dead'" class="cell"
      ></td>
    </tr>
  </table>
</template>

<script>

import axios from 'axios'

export default {
  name: 'App',
  components: {
  },

  data() {
    return{
      size: 0,
      probability: 0.0,
      cells: [],
      //grid: []
    }
  },

  methods:{
    setDimensions(size){
      this.size = size
    },
    setGrid() {
      this.cells = []
      // this.grid = []
      for(let i=0; i<this.size; i++){
        this.cells[i] = []
        for (let j = 0; j < this.size; j++) {
          if(this.randomNumber() > (this.probability)*10){
            this.cells[i][j] = false 
            //this.grid[i][j]= "_"
          } 
          else{
            this.cells[i][j] = true
            // this.grid[i][j]= "*"
          } 
        }
      }
      //console.log(this.cells)
      this.postreq()
    },
    changeState(cell, i,j) {
      this.cells[i][j] = !cell
      //if (this.grid)
      //else 
      //console.log(cell)
    },
    randomNumber() {
      return Math.floor(Math.random() * (10 - 1 + 1)) + 1
    },
    postreq: function() {
      var data = {"grid": this.cells}
      /*eslint-disable*/
      console.log(data) 

      axios({ method: "POST", url: "http://127.0.0.1:3000/game", data: data, headers: {"content-type": "text/plain" } }).then(result => { 
          // this.response = result.data;
          /*eslint-disable*/
          console.log(result.data) 
          /*eslint-enable*/
        }).catch( error => {
            /*eslint-disable*/
            console.error(error);
            /*eslint-enable*/
      });
    }

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
  margin: 1rem 0.25rem;
}

.cell {
  padding: 0.26rem;
}

.alive {
  border-style:solid;
  background-color: black ;
  border-width: 1px;
}

.dead {
  border-style:solid;
  background-color: white ;
  border-width: 1px;
}

table {
  margin: auto;
}

</style>