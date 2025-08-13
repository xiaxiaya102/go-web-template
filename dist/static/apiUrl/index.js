import apc from "./apc";
import file from "./file";
import theme from "./theme";
import user from "@static/apiUrl/user";

export default {
  ...apc,
  ...file,
  ...theme,
  ...user,
};
