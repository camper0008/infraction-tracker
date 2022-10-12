const makeOptions = () => {
    const apiKey = document.getElementById("api-key");
    const headers = new Headers({ "X-API-KEY": apiKey.value });
    const method = "POST";

    const options = { headers, method };
    return options;
};

const remove = async (name) => {
    const options = makeOptions();
    const res = await (await fetch(`/api/remove/${name}`, options)).json();
};

const update = async (name) => {
    const count = document.getElementById("count-" + name);
    const options = makeOptions();

    const res = await (
        await fetch(`/api/update/${name}/${count.value}`, options)
    ).json();
};

const add = async () => {
    const name = document.getElementById("create-name");
    const options = makeOptions();

    const res = await (await fetch(`/api/add/${name.value}`, options)).json();
};
