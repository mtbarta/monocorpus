//somehow this is the only moment import that has a problem.
let moment = require("moment");
if ("default" in moment) {
    moment = moment["default"];
}


const PRETTY_PRINT = 'YYYY-MM-DD'
const PRETTY_PRINT_WITH_TIME = 'YYYY-MM-DD hh:mm A'

export const normalizeDate = (timestamp: number, withTime?: boolean) : string => {
  if (!withTime) {
    return moment.unix(timestamp).local().format(PRETTY_PRINT)
  }
  
  return moment.unix(timestamp).local().format(PRETTY_PRINT_WITH_TIME)
}

export const today = moment().local().format(PRETTY_PRINT)