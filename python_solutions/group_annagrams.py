def group_annagrams(li: list) -> list:
    words_list_mapping = {}
    for word in li:
        char_li = list(word)
        char_li.sort()
        char_li = tuple(char_li)
        if char_li in words_list_mapping:
            words_list_mapping[char_li].append(word)
        else:
            words_list_mapping.update({char_li: [word]})

    return list(words_list_mapping.values())


print(group_annagrams(["asdas", "asdas", "asdas", "asdas"]))
