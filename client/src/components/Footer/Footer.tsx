import React from "react"
import { getCurrentYear } from "../../utils/helpers"

const Footer: React.FC = () => {
    return (
        <footer className="bg-white my-4 text-zinc-600 p-4 relative bottom-0">
          <div className="container mx-auto text-center">
            <p>&copy; {getCurrentYear()} XYZ Books</p>
            <p>Created by Engr. Kirk Alyn Santos</p>
          </div>
        </footer>
      )
}

export default Footer