import NavBar from './components/NavBar/NavBar'
import Footer from './components/Footer/Footer'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"
import Books from './components/Books/Books'
import Authors from './components/Authors/Authors'
import Publishers from './components/Publishers/Publishers'
import BookDetails from './components/BookDetails/BookDetails'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Books />,
  },
  {
    path: "/books/:isbn",
    element: <BookDetails />,
  },
  {
    path: "/authors",
    element: <Authors />,
  },
  {
    path: "/publishers",
    element: <Publishers />,
},
]);

function App() {
  return (
    <main>
      <NavBar />
        <RouterProvider router={router} />
      <Footer />
    </main>
  )
}

export default App
