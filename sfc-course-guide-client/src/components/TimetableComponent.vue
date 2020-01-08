<template>
  <div class="timetable" @click="appearTable">
    <!-- @click="loadTable" v-if="showTable" @close="closeTable" -->
    <table id="test" border="1">
      <thead>
        <tr>
          <th style="width: 100px;"></th>
          <th v-for="day in days" style="width: 200px;">{{day}}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="n in [0, 1, 2, 3, 4, 5]">
          <th>{{n}}</th>
          <td v-for="v in [0, 1, 2, 3, 4, 5]">&nbsp;</td>
        </tr>
      </tbody>
    </table>
    <button @click="reset_timetable">Reset</button>
  </div>
</template>

<style scoped>
.timetable {
  /* padding: 50px;
  width: 1000px;
  height: 800px; */
  text-align: center;
  background-color: rgb(192, 147, 79);
  /* z-index: 10px; */
  /* position: absolute; */
  top: 0;
  bottom: 0;
  /* left: -970px; */
  /* margin-left: auto; */
  /* margin-right: auto; */
  border-radius: 10%;
}

/* .timetable:hover {
  transform: translate(1000px, 0px);
} */

table {
  color: rgb(75, 75, 75);
  /* border: 0.5px solid #aaa;
  border-collapse: separate;
  border-spacing: 0;
  border-radius: 6px; */
  overflow: hidden;
  background-color: rgb(255, 255, 255);
  /* width: 800px; */
  /* margin-left: auto;
  margin-right: auto; */
  border: none;
}
table thead th,
table tbody th,
table tbody td {
  /* padding: 0.6em 3em; */
  /* border-bottom: 0.5px solid #aaa; */
  border: none;
}
table thead th {
  background-color: rgb(79, 192, 141);
}
table tbody th {
  background-color: #eee;
}
table thead th + th,
table tbody td {
  border-left: 0.1px solid #aaa;
  /* border: none; */
}
table tbody tr:last-child th,
table tbody tr:last-child td {
  border-bottom: none;
}

body.dark table {
  background-color: #282c34;
  color: rgb(221, 221, 221);
}

body.dark table thead th {
  color: #282c34;
  background-color: rgb(192, 147, 79);
}

body.dark table tbody th {
  color: #282c34;
}
</style>

<script>
export default {
  data: function() {
    return {
      days: ["月", "火", "水", "木", "金", "土"],
      showTable: true
    };
  },
  mounted() {
    if (localStorage.local_array) {
      var table = document.getElementById("test");
      var array = JSON.parse(localStorage.local_array);
      for (var i = 0; i <= array.length; i++) {
        var titles_en = array[i]["title"];
        var x = array[i]["x"];
        var y = array[i]["y"];
        var td = table.rows[y].cells[x];
        td.firstChild.nodeValue = titles_en;
      }
    }
  },
  methods: {
    reset_timetable() {
      localStorage.clear();
      location.reload();
    },
    appearTable() {
      this.showTable = true;
    },
    closeTable() {
      this.showTable = false;
    }
  }
};
</script>