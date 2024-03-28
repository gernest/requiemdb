import { createRoutesFromElements, Route } from "react-router-dom";
import { NotFound, RouterDataErrorBoundary, Layout, Console } from "./pages";

export const routes = createRoutesFromElements(
    <Route element={(<Layout />)} errorElement={<RouterDataErrorBoundary />}>
        <Route path="*" element={<NotFound />} />
        <Route path="console" element={<Console />} />
    </Route>
)