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
      showModal: false
    };
  },
  props: ["searchResult", "query"],
  computed: {
    course() {
      return this.searchResult;
    }
  },
  methods: {
    loadModal() {
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
    },
    add_to_timetable() {
      const titles_en = this.course.title;

      const days_en = this.course.class_day;

      const days_list = days_en.split(",");
      console.log(days_list[0]);
      let x;
      let y;

      const week_list = [
        "月曜日",
        "火曜日",
        "水曜日",
        "木曜日",
        "金曜日",
        "土曜日"
      ];

      console.log(week_list);
      for (let i = 0; i < week_list.length; i++) {
        if (days_list[0].indexOf(week_list[i]) !== -1) {
          x = i + 1;
        }
      }
      console.log(x);

      const numbers = [
        "１時限",
        "２時限",
        "３時限",
        "４時限",
        "５時限",
        "６時限"
      ];

      for (let i = 0; i < numbers.length; i++) {
        if (days_list[0].indexOf(numbers[i]) !== -1) {
          y = i + 1;
        }
      }
      console.log(y);

      const table = document.getElementById("test");
      const td = table.rows[y].cells[x];
      td.firstChild.nodeValue = titles_en;

      const array = {};
      const top_array = [];
      array["title"] = titles_en;
      array["x"] = x;
      array["y"] = y;
      if (localStorage.local_array) {
        const top_array = JSON.parse(localStorage.local_array);
      } else {
        const top_array = [];
      }
      top_array.push(array);
      // converting the array into a string
      localStorage.local_array = JSON.stringify(top_array);
      console.log(top_array);
    }
  }
};
</script>
