import axios from "axios"

export const authPin = async ({ pin, name }: { pin: string; name: string }) => {
  return await axios.post("http://localhost:8080/api/auth-pin", {
    name,
    pin,
  })
}
