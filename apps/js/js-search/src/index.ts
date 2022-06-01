import movies from "../../../../data/movies.json";

const anyArray = movies as any[];

const searchInObject = (obj: any, searchString: string) => {
  for (const key in obj) {
    if (obj.hasOwnProperty(key)) {
      const value = obj[key];

      if (typeof value === "object") {
        const searchResult = searchInObject(value, searchString);
        if (searchResult) {
          return searchResult;
        }
      } else if (typeof value === "string") {
        if (value.toLowerCase().includes(searchString.toLowerCase())) {
          return obj;
        }
      }
    }
  }
};

const search = (datas: any[], searchString: string) => {
  const searchResults = [] as any[];

  for (const data of datas) {
    const searchResult = searchInObject(data, searchString);
    if (searchResult) {
      searchResults.push(searchResult);
    }
  }
  return searchResults;
};

const startTime = Date.now();
const values = search(anyArray, "Science Fiction");
const endTime = Date.now();
const elapsedTime = endTime - startTime;
console.log(values);
console.log(`Found ${values.length} results in ${elapsedTime}ms`);
