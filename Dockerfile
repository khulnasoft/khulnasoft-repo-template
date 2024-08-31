FROM ubuntu:latest
SHELL ["/bin/bash", "-c"]
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
    apt-get upgrade --yes && \
    apt-get install --yes software-properties-common
RUN add-apt-repository ppa:neovim-ppa/unstable 

RUN apt-get update && \
    apt-get install --yes \
    docker.io \
    wget \
    ripgrep \
    xclip \
    ca-certificates \
    docker-compose-v2 \
    neovim \
    curl \
    openssh-server \
    tmux \
    zsh \
    fzf \
    locales \
    openssl \
    man \
    sudo \
    vim \
    unzip \
    build-essential \
    git \
    && rm -rf /var/lib/apt/lists/*

RUN sed -i '/en_US.UTF-8/s/^# //g' /etc/locale.gen && locale-gen
ENV LANG en_US.UTF-8  
ENV LANGUAGE en_US:en  
ENV LC_ALL en_US.UTF-8

ARG USER
ARG PASSWORD
ARG ZSHRC_FILE_URL
ARG P10K_FILE_URL
ARG NEOVIM_CONFIG_REPO_URL
ARG EXPOSE_INTERNAL_PORT_RANGE
ARG EXPOSE_INTERNAL_PATH_DIRECTORY

RUN useradd ${USER} \
    --create-home \
    --shell=/usr/bin/zsh \
    --groups=sudo,docker \
    --uid=1001 \
    --password "$(openssl passwd -6 $PASSWORD)" \
    --user-group && \
    echo "${USER} ALL=(ALL) NOPASSWD:ALL" >>/etc/sudoers.d/nopasswd

USER ${USER}
WORKDIR /home/${USER}
SHELL ["/usr/bin/zsh", "-c"]

RUN wget https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh -O - | zsh || true
RUN git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k
RUN git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
RUN git clone https://github.com/Aloxaf/fzf-tab ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/fzf-tab
RUN git clone https://github.com/zdharma-continuum/fast-syntax-highlighting.git \
  ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/plugins/fast-syntax-highlighting
RUN rm .zshrc && mkdir codes && mkdir -p ${EXPOSE_INTERNAL_PATH_DIRECTORY}
RUN wget ${P10K_FILE_URL} -O .p10k.zsh \
    && wget ${ZSHRC_FILE_URL} -O .zshrc
RUN git clone ${NEOVIM_CONFIG_REPO_URL} /home/${USER}/.config/nvim
COPY devenvironment.sh .
RUN sudo chmod +x devenvironment.sh && ./devenvironment.sh
EXPOSE 22/tcp
EXPOSE ${EXPOSE_INTERNAL_PORT_RANGE}
RUN source ~/.zshrc
RUN sudo rm -rf .zcompdump* .wget-hsts .sudo_as_admin_successful .bashrc .bash_logout .bash_history devenvironment.sh
ENTRYPOINT sudo service ssh start && echo "$(service ssh status)" && nohup sudo dockerd >/dev/null 2>&1 && tail -f /dev/null
