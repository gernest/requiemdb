import { Box, Text, Label, Link } from "@primer/react"
import { MarkGithubIcon } from "@primer/octicons-react"
import { useEffect, useState } from "react";
import { useApi } from "../../providers";


export const Footer = () => {
    const [version, setVersion] = useState<string>("unknown")
    const { client } = useApi();
    useEffect(() => {
        client.getVersion({}).then(build => {
            setVersion(build.response.version)
        }).catch(e => {
            console.log(e)
        })
    }, [client])
    return (
        <Box
            position={"fixed"}
            zIndex={9999}
            bottom={0}
            sx={{
                display: "flex",
                height: "45px",
                backgroundColor: 'canvas.subtle',
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
                alignItems: "center",
            }}
            >
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