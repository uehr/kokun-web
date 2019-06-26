const apiEndpoint = "https://kokun.herokuapp.com/senryu"

function getSenryuImage(firstSentence, secondSentence, thirdSentence, authorName) {
    const postSenryu = {
        "FirstSentence": firstSentence,
        "SecondSentence": secondSentence,
        "ThirdSentence": thirdSentence,
        "AuthorName": authorName
    }

    return postData(apiEndpoint, postSenryu)
}

function postData(url = ``, data = {}) {
    return fetch(url, {
        method: "POST",
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "application/json; charset=utf-8",
        },
        redirect: "follow",
        referrer: "no-referrer",
        body: JSON.stringify(data),
    })
        .then(response => response.json());
}