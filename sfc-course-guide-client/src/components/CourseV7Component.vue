<template>
  <div class="course-box" @click="loadModal">
    <div class="course-info">
      <div class="subject-sort" v-html="course.kamoku_sort"></div>
      <div class="title-memo-en" v-html="course.title_memo"></div>
      <!-- <div class="language-en" v-html="course.Language_en"></div> -->
    </div>
    <div class="title-en" v-html="course.title"></div>
    <div class="semester-en" v-html="course.semester"></div>
    <div class="days-en" v-html="course.class_day"></div>
    <!-- <div class="faculty-in-charge" v-html="course.Faculty_in_charge"></div> -->
    <div class="add-to-timetable">
      <button @click="add_to_timetable">Add</button>
    </div>
  </div>
</template>

<style>
.course-box {
  border-radius: 25px;
  padding: 20px;
  width: 25em;
  margin: 1%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  line-height: 2em;
  box-shadow: 0px 0px 4px #c4c4c4;
}

.course-box:hover {
  cursor: pointer;
  background: Highlight;
}

.course-box .highlight {
  color: red;
}

.course-info {
  display: inline-flex;
  justify-content: space-evenly;
}

.subject-sort {
  float: left;
}

.language-en {
  float: right;
}

.title-en {
  color: rgb(79, 192, 141);
  font-weight: bold;
}

body.dark .title-en {
  color: rgb(192, 147, 79);
  font-weight: bold;
}
</style>
<style class="dark">
body.dark .course-box {
  border-radius: 25px;
  padding: 20px;
  width: 25em;
  margin: 1%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  line-height: 2em;
  box-shadow: 0px 0px 4px #c4c4c4;
}

body.dark .course-box .highlight {
  color: #ff6200;
}
</style>
<script>
export default {
  data() {
    return {
      showModal: false,
    };
  },
  props: ['searchResult', 'query'],
  computed: {
    course() {
      return this.searchResult;
    },
  },
  methods: {
    loadModal() {
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
    },
    add_to_timetable() {
      var titles_en = this.course.Title_en.replace(
        /<("[^"]*"|'[^']*'|[^'">])*>/g,
        ""
      );

      var days_en = this.course.Days_en.replace(
        /<("[^"]*"|'[^']*'|[^'">])*>/g,
        ""
      );

      var days_list = days_en.split(" ");

      var x;
      var y;

      switch (days_list[0]) {
        case "Monday":
          x = 1;
          break;
        case "Tuesday":
          x = 2;
          break;
        case "Wednesday":
          x = 3;
          break;
        case "Thursday":
          x = 4;
          break;
        case "Friday":
          x = 5;
          break;
        case "Saturday":
          x = 6;
          break;
      }

      switch (days_list[1]) {
        case "1st":
          y = 1;
          break;
        case "2nd":
          y = 2;
          break;
        case "3rd":
          y = 3;
          break;
        case "4th":
          y = 4;
          break;
        case "5th":
          y = 5;
          break;
        case "6th":
          y = 6;
          break;
      }

      var table = document.getElementById("test");
      var td = table.rows[y].cells[x];
      td.firstChild.nodeValue = titles_en;

      var array = {};

      array["title"] = titles_en;
      array["x"] = x;
      array["y"] = y;
      if (localStorage.local_array) {
        var top_array = JSON.parse(localStorage.local_array);
      } else {
        var top_array = [];
      }
      top_array.push(array);
      // converting the array into a string
      localStorage.local_array = JSON.stringify(top_array);
    }
  },
};
</script>
