import { Box, Text, Label, Link } from "@primer/react"
import { MarkGithubIcon } from "@primer/octicons-react"
import { useState } from "react";


export const Footer = () => {
    const [version, setVersion] = useState<string>("unknown")
    return (
        <Box
            sx={{
                display: "flex",
                height: "45px",
                backgroundColor: 'canvas.subtle',
                zIndex: 2000,
            }}
            width={"100%"}
        >
            <Box
                sx={{
                    display: "flex",
                    paddingLeft: "1rem",
                    alignItems: "center",
                    flex: "1",
                }}
            >
                <Text>
                    Copyright &copy; {new Date().getFullYear()} Geofrey Ernest
                </Text>
            </Box>
            <Box sx={{
                display: "flex",
                paddingRight: "1rem",
                alignItems: "center",
            }}>
                <Label variant="primary" sx={{
                    marginRight: 1,
                }}>
                    requiem: {version}
                </Label>
                <Link
                    href='https://github.com/gernest/requiem'
                    target='_blank'
                    rel='noreferrer'
                >
                    <MarkGithubIcon size={"medium"} />
                </Link>
            </Box>
        </Box>
    )
}